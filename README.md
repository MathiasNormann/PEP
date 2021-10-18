# PEP

The solution is made using golang, run in a docker container.

Created by
```
docker build -t "pep" .
```
And run by
```
docker run -p 3000:3000 pep
```
And the tests are then run by running
```
./runTests.sh
```

First test is {"Mette Frederiksen", 19.11.1977} and second test is {"Hans Christian Andersen", 02.04.1805}

The solution simply reads a csv file, collects all persons and looks up a person given as command arguments to the solution to check whether the person exists in the list.

How to procede with future development as a standalone solution:
- ~Make it a service that does not need to read file for every look-up~
- Set up a DB to store all PEP in
- Add update functionality to the DB

Regarding using a Graph DB as the DB for a solution there a different possibilities:
- Define own ontology of PEP (easy to do and manage, but cannot easily be used together with other ontologies)
- Use FoaF for a PEP (already defined and is the most used ontology for persons)
If we consider where they work it gets more interesting. There we have for example the option to use the euBusinessGraph to hopefully align with other ontologies used inside the EU.
