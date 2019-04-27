<h1> System For Gathering Censorship Measurement Data </h1>
This is our project for CMPSCI 660 Advanced Information Assurance!<
We are attempting to redesign the system discussed in:https://conferences.sigcomm.org/sigcomm/2015/pdf/papers/p653.pdf
The system proposed was critisized as being unethical due to the users being unaware of the deployment of the system which would automatically perform the measurement tasks (which could potentially land them in hot water with the censor). We attempt to redesign the system in such a way that the users ARE aware of the system and can opt-in, while maintaining the automated form of data coalition for censored sites and where they are blocked.

To run our system:<br>
For the server:
1. install required libraries
2. run the command: go run server.go database.go testingGeoIP.go

<br>For the user:<br>
1. Install greaseMonkey or Tampermonkey (depending on browser)
2. install the script contained in this github repository named: censorshipMeasurementScript.js

```bash
├── README.md
├── censorshipMeasurementScript.js
├── cert
│   ├── server.crt
│   └── server.key
├── data
│   ├── TaskCreator.java
│   ├── URLUtils.java
│   ├── expandedTest.txt
│   └── shortgfwlist.txt
├── database.go
├── server.go
├── task_templates
│   ├── Procfile
│   └── img.js
├── testingGeoIP.go
└── webpage
    ├── index.html
    └── ourCSS.css
```



