<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
This project has been created as a master's thesis at the University of Coruña. 
```
The dairy sector seeks to offer mechanisms with a quality certification of the products produced in our farms, both at the level of quality standards of the raw material, as well as the production characteristics or the ethical context associated with it. The importation of milk has forced the dairy industry to look for different methods to give added value to the product to be sold. This master’s thesis aims to offer, through the latest traceability technologies, a seal of quality, providing clear and transparent information on the origin and the different locations and processes through which the material passes until it leaves the processing plant. The tool to be developed contemplate in its design the entire dairy production chain, although it focus essentially on traceability associated with three data sources. The first of these is related to the storage of milk by the farmer in the relevant storage silos. The second source of traceability data is associated with the collection of the raw material by the transporter, who takes the milk from origin to the processing plant. The last trace is collected at the plant, where the milk is processed. Each of the aforementioned data sources result in a record where different information stored. The blockchain-based system store data of interest related to the most
important aspects in the whole dairy chain.
```

**Table of Contents**  *generated with [DocToc](https://github.com/thlorenz/doctoc)*

- [Blockchain restful API](#golang-template-project)
  - [About the project](#about-the-project)
    - [API docs](#api-docs)
  - [Getting started](#getting-started)
    - [Layout](#layout)
  - [Notes](#notes)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

# Blockchain restful API

## About the project

This project is an API to make request against a blockchain using a smart contract.

### API docs

* AddNewTransport: POST petition what creates a new transport. 
* GetTransportById: GET petition to get a transport by id.
* TransportRouteAddFarmRecollectionToTransport: POST petition to add a farm in a collection.
* TransportRoutePopFarmRecollectionToTransport: POST petition to remove a farm in a collection.
* GetTraces: GET petition to get all traces.
* GetTraceById: GET petition to get all traces filter by id.
* AddNewTrace: POST petition to create a new trace.
* TraceRoutAddFarmToTrace: POST petition to add a farm in a trace.
* TraceRoutAddTransvaseToTrace: POST petition to add a transvase in a trace.

## Getting started
First of all, need to set up a blockchain. If you want to use test-network of hyperledger tou can execute:
```
./network.sh up createChannel -c mychannel -ca
./network.sh deployCC -ccn basic -ccp ../asset-transfer-basic/chaincode-go/ -ccl go
```
After this set up the server:

```
$ ~/go/bin/wire
$ ./blockchain-restful-api
```
Below we describe the conventions or tools specific to golang project.

### Layout

```tree
├── .github
├── .gitignore
├── README.md
├── api
│   ├── commom
│   │   └── constants.go
│   └── controllers
│   |    └── trace.go
│   |    └── transport.go
|   └── repositories
|   |   └── trace.go
|   |   └── transport.go
│   └── api.go
├── blockchain
│   └── erros.go
│   └── hyperledger-app.go
│   └── interface.go
├── models
│   ├── models.go
├── main.go
├── wire.go
├── wire_gen.go

```


## Notes
