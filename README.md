# Log


## 15.12.2023: Lezione 0
https://www.youtube.com/watch?v=TtCfDXfSw_0&list=PLy_6D98if3ULEtXtNSY_2qN21VCKgoQAE&ab_channel=TECHSCHOOL

Ho installato sqlc con `go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest`

## 15.12.2023: Lezione 1
https://www.youtube.com/watch?v=Q9ipbLeqmQo&list=PLy_6D98if3ULEtXtNSY_2qN21VCKgoQAE&index=3&ab_channel=TECHSCHOOL


## 15.12.2023: Lezione 2


## 17.12.2023: Lezione 3
https://www.youtube.com/watch?v=0CYkrGIJkpw&list=PLy_6D98if3ULEtXtNSY_2qN21VCKgoQAE&index=4

Schema migration 
usiamo golang migrate library per schema migration 

installato go migrate cli via brew

poi ho creato un target in makefile per fare la migrazione. 

`make migrate` fa il trick. 

Istruzioni qui: https://github.com/golang-migrate/migrate/blob/master/GETTING_STARTED.md#create-migrations

Importante da notare:
> IMPORTANT: In a project developed by more than one person there is a chance of migrations inconsistency - e.g. two developers can create conflicting migrations, and the developer that created their migration later gets it merged to the repository first. Developers and Teams should keep an eye on such cases (especially during code review). Here is the issue summary if you would like to read more.

> Consider making your migrations idempotent - we can run the same sql code twice in a row with the same result. This makes our migrations more robust. On the other hand, it causes slightly less control over database schema - e.g. let's say you forgot to drop the table in down migration. You run down migration - the table is still there. When you run up migration again - CREATE TABLE would return an error, helping you find an issue in down migration, while CREATE TABLE IF NOT EXISTS would not. Use those conditions wisely.


L'idea e':
- prima faccio il design del db in [dbdiagram] (https://dbdiagram.io/d/Simple-bank-63b48c837d39e42284e8b75c) 
- poi esporto il codice postgressql per creare quel db
- creo una migrazione 
- copio tutto questo codice generato direttamente nei migration files

## 18.12.2023: Lezione 4 - CRUD operations sul db
https://www.youtube.com/watch?v=prh0hTyI1sU&list=PLy_6D98if3ULEtXtNSY_2qN21VCKgoQAE&index=5

In pratica ho installato sqlc... 

Con `sqlc init` mi crea un file sqlc.yaml in root.

A questo punto sono andato sul sito di sqlc ed ho copiato un config di esempio... a cui ho sostituito i valori come nel tutorial.
Attenzione, io sto usando la versione 2, nel tutorial usa la versione 1. Bisogna adattarla campo per campo. 
Nel mio caso ho levao il field cloud e il riferimento a managed db.

Ricapitolando:
- in `db` folder abbiamo:
    - migration: contiene tutte le migrazioni. Vi punto anche dal sqlc.yaml
    - query: e' la folder dove inserisco tutte le query che do in input a sqlc... contenenti le varie annotazioni di sqlc
    - sqlc: e' la mia folder ti output contenenti tutto il codice generato

In pratica ho scritto tutte le query per crud in account.sql e generato il codice corrispondente. 

19.12.2023: Lezione 5: write unit test for database CRUD
https://www.youtube.com/watch?v=phHDfOHB2PU&list=PLy_6D98if3ULEtXtNSY_2qN21VCKgoQAE&index=6

Installato il driver pg: `go get github.com/lib/pq`

Poi abbiamo creato i test... main_test.go e account_test.go


NB. ho dovuto commentare la riga `sql_package: "pgx/v5"` in sqlc.yaml altrimenti la conn non funzionava... non veniva presa salla New()

Ho anche creato il random generator e un nuovo target per test nel Makefile. 

## 19.12.2023: Lezione 6 - A clean way to implement database transaction in Golang
https://www.youtube.com/watch?v=gBh__1eFwVI 

Qui si introducono le transactions.

Come prima cosa creo un nuovo file store.go in cui definisco una struct in cui definire il supporto per le transazioni. 

Questa struct fornisce tutte le funzioni per eseguire le query individualmente cosi come la loro combinazione con le transazioni.

Per le query individuali abbiamo gia la struct Queries generata da sqlc. 

Each query only does one query on a specific table. Queries struct does not support transactions. 
This is why we have to extend its functionality by embedding inside the struct by composition.

We need to embed also the db object in order to support for transactions 


`execTx` e' una funzione che serve per eseguire una transazione generica. L'idea e' semplice. Il concetto e' prendere un context e una callback function come input ... una sorta di wrapper. 

BeginTx e' l'inizio della transazione... 

Allochiamo un oggetto di tipo `Queries` che mi viene restituito dalla New generata dal `sqlc`.

NB. il parametro options passato a BeginTx serve per settare l'isolation method.

La struttura e' una sorta di wrapper con:
- begin
- rollback in caso di errore
- commit in caso di successo


## 19.12.2023: Lezione 7 - DB transaction lock & How to handle deadlock in Golang
https://www.youtube.com/watch?v=G2aggv_3Bbg

Qui vediamo i transaction locks.

Usa TDD per questo video.

In soldoni, fa notare come transfer tx contiene un deadlock... 
e questo deadlock e' dovuto al fatto che viene richiesto un lock su una select dovuta alle foreign keys.
C'é una dipendenza tra account e transfer table via foreign key... 
Se la si disabilita ... si risolve il deadlock... 
ma non e'la soluzione ideale... vogliamo mantenere la foreign key contraint...
quindi lo rimette ... e prova a risolvere in un altro modo.

.... 

Come risolve? Aggiungendo "FOR NO KEY UPDATE"in GetAccountForUpdate. In tal modo diciamo a postgres di non fare lock per quella select in quanto non andremo a toccare quella chiave... 

Alla fine fa vedere come migliorare il codice della transferTx... aggiungendo AddAccountBalance.

NB. per modificare e rinominare un argument in sqlc... 
invece di $2, usa sqlc.arg(NEWNAME)

NEWNAME e' il nome del nuovo parametro generato nella corrispondente request struct. 


## 22.12.2023 - Lezione 8 - how to avoid db deadlock

Uso questa query per controllare i lock:
```sql
SELECT a.application_name,
         l.relation::regclass,
         l.transactionid,
         l.mode,
		 l.locktype,
         l.GRANTED,
		 a.pid,
         a.usename,
         a.query,
         a.query_start,
         age(now(), a.query_start) AS "age",
         a.pid
FROM pg_stat_activity a
JOIN pg_locks l ON l.pid = a.pid
WHERE a.application_name = 'psql'
ORDER BY a.query_start;
```

Da questa query vedo quali query hanno un lock, lo stato del lock (granted, cioe' se e' stato concesso o no) e tante altre info.

Per esempio, la transaction id, il tipo di lock, se shared o exclusive. 

Piu' dettagli qui... nb. transaction id mi dice l'id della transazione che sta aspettando la transazione corrente.

(ID of the transaction targeted by the lock, or null if the target is not a transaction ID)

https://www.postgresql.org/docs/current/monitoring-stats.html

How to prevent deadlocks? 
To make sure all applications always aquire locks in a consistent order. 

Che significa? 
Se una transazione deve fare l'update di due account, account1 e account2.
Possiamo definire una regola secondo la quale facciamo l'update a partire dall'account id piu' piccolo. 

In questa lezione abbiamo anche creato un metodo helper AddMoney per refactoring.



## 26.12.2023 - Lezione 9 - Transaction isolation level
https://www.youtube.com/watch?v=4EajrPgJAk0

Spiega prima la teoria poi fa qualche esempio con pqsl per dimostrare i vari tipi di isolation.


## 26.12.2023 - Lezione 10 - Setup Github Actions for Golang + Postgres to run automated tests
https://www.youtube.com/watch?v=3mzQRJY1GVE

Ho creato un nuovo repository in github e pushato il codice. 
L'ho chiamato work_simplebank (a differenza di quello esistente work-simplebank).




# Appendix

## DBTX interface

Viene generata da SQLC.

```go
type DBTX interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}
```

Un oggetto "db" che implementa questa inferface puo' essere passato alla factory `New`

NB. La struct `database/sql` DB implementa l'interfaccia `DBTX` 

motivo per cui possiamo passarlo alla factory `New` senza problemi.





## PGX: 

**pgx** is a pure Go driver and toolkit for PostgreSQL. It’s become the default PostgreSQL package for many Gophers since lib/pq was put into maintenance mode.

To start generating code that uses pgx, set the sql_package field in your sqlc.yaml configuration file. Valid options are pgx/v4 or pgx/v5

