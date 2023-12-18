# Log


15.12.2023: Lezione 0
https://www.youtube.com/watch?v=TtCfDXfSw_0&list=PLy_6D98if3ULEtXtNSY_2qN21VCKgoQAE&ab_channel=TECHSCHOOL

Ho installato sqlc con `go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest`

15.12.2023: Lezione 1
https://www.youtube.com/watch?v=Q9ipbLeqmQo&list=PLy_6D98if3ULEtXtNSY_2qN21VCKgoQAE&index=3&ab_channel=TECHSCHOOL


15.12.2023: Lezione 2


17.12.2023: Lezione 3
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

18.12.2023: Lezione 4 - CRUD operations sul db
https://www.youtube.com/watch?v=prh0hTyI1sU&list=PLy_6D98if3ULEtXtNSY_2qN21VCKgoQAE&index=5
