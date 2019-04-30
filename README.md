<h1> System For Gathering Censorship Measurement Data </h1>
This is our project for CMPSCI 660 Advanced Information Assurance!<br>
We are attempting to redesign the system discussed in:https://conferences.sigcomm.org/sigcomm/2015/pdf/papers/p653.pdf
The system proposed was critisized as being unethical due to the users being unaware (and unable to opt-out of) the use of thier machine to perform measurement tasks (which might potentially land them in hot water with the censor). We attempt to redesign the system in such a way that the users ARE aware of the system and can opt-in, while maintaining the automated form of data coalition for censored sites and where they are blocked.

To run our system:<br>
For the server:
1. install go and mySQL
2. install required libraries (go-sql-driver/mysql)
3. ensure mysql service is running (computer (right click)-> select "manage" -> go to "services" -> find "mysql" (right click) -> select start)
4. If you wish to have different measurment tasks call the database.addTask(url, target) to include them in the database.
5. Update line 26 of database.go to reflect the username password and database name for where you wish to store the results. Form: "(username):(password)@tcp(localhost)/(database_name)")
6. run the command: go run server.go database.go testingGeoIP.go


<br>For the user:<br>
1. Install greaseMonkey or Tampermonkey (depending on browser)
2. update line 31 of the greasemonkey script (censorshipMeasurementScript.js) to point to the address where the server is running.
3. install the script contained in this github repository named: censorshipMeasurementScript.js

```bash
├── README.md
├── censorshipMeasurementScript.js          #greesemonkey script
├── cert
│   ├── server.crt
│   └── server.key
├── data                                    #tools for creating tasks in our database
│   ├── TaskCreator.java
│   ├── URLUtils.java
│   ├── expandedTest.txt
│   └── shortgfwlist.txt
├── database.go                             #the backend class that handles database interactions
├── server.go                               #the backend class that handles rest requests from clients
├── task_templates
│   ├── Procfile
│   └── img.js                              #the template for a task 
├── testingGeoIP.go                         #class used for taking client ip and identifying where they are geographically located
└── webpage                                 #for demo purposes on how the system works
    ├── index.html      
    └── ourCSS.css
```



