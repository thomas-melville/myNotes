# nfv

Network Function Virtualization

Founded by 7 of the worlds leading telecoms network operators.
the home of the Industry Specification Group for NFV

The purpose of the group is to build the Software Defined Network, SDN
Which will reduce capex, opex and time to market
NFV and SDN are complementary but increasing co-dependent

NFV Release 3 is under way, 2017 & 2018
Releases are in 2 year phases

## MANO

Management And Organization
Managing NFV
Working Group within ETSI

Many is broken up into 3 parts

NFVO
VNFM
VIM

(They're detailed below)

## SOL

Is a new working group in ETSI NFV, SOLutions
Charged with the task of delivering a consolidated set of protocols and data model specs to support interoperability

## IFA

is another working group
Interfaces And Architecture

## Network Service

NS
A bunch of VNFs with connective and service function chaining

## NFVO

Network Function Virtualization Orchestrator
Manages the lifecycle of Network Services
Communicates with OSS/BSS

ECM would be an NFVO, correct

## VIM

Virtualized Infrastructure Manager
Manages the NFVI resources: compute, network and storage
The clouds itself, example openstack / kubernetes
Communicates with the VNFM

## NFVI

Network Function Virtualization Infrastructure
compute, network and storage

## VNFM

Virtual Network Function Manager
Manages the lifecycle of VNFs

## Reference Points

Interfaces between different "actors" in the NFV space.

Actors: VNF, VNFM, VIM, NFVO

HTTP protocol
Each reference point defines multiple interfaces between the components which the reference point links.
in both directions with HTTP Methods, query parameters, body and headers specified

### URI structure

{apiRoot}/{apiName}/{apiVersion}/

the body will be in JSON format

HTTPS, TLS 1.2

#### apiRoot

protocol + host + (port) + (prefix)

#### apiName

interface name in the abbreviated form

#### apiVersion

current version of the API

### attribute based filtering

attribute based filtering is used on GET requests to filter whats returned.
attributes of the object can be compared, and attributes deeper in the structured document can be retrieved for comparison using .
operator can be text based and nested off attribute to be compared too.

ex: GET .../container?weight.eq=100 or GET .../container?weight=100

operators: eq, neq, gt, lt, gte, lte, cont, ncont

### attribute selectors

allow the API consumer to decide which attributes it wants contained in the response.
attributes can be marked for inclusion or exclusion
for omitted attributes, a link to a resource may be returned where the info can be fetched from. HATEOAS
These links will be in a _links object at the same level as the omitted attribute

query parameters:

* all_fields        include all fields
* fields            comma seperated list of attributes, same as above for attributes deeper in the structured doc
* exclude_fields
* exclude_default

### Error reporting

app errors are mapped to HTTP Errors. 4xx / 5xx
body should contain json representation of a ProblemDetails structure

* type      URI to definition of problem type
* title     short, human readable summary of problem type
* status    HTTP status code
* detail    explanation specific to this occurrence
* instance  URI reference that identifies the specific occurrence
* other     any other additional attributes you want to add in

### or-vnfm

Reference point which contains the interfaces between the VNFM (Manager) and the NFVO (Orchestrator)

**Interfaces.**

Component on the left produces the interface, component on the right consumes the interface

* VNF Lifecycle Management Interface. VNFM -> NFVO
* VNF Performance Management Interface. VNFM -> NFVO
* VNF Fault Management Interface. VNFM -> NFVO
* VNF Indicator Interface. VNFM -> NFVO
* VNF Lifecycle Operation Granting Interface. NFVO -> VNFM
* VNF Package Management Interface. NFVO -> VNFM
* Virtualized Resources Quota Available Notification Interface. NVFO -> VNFM

### vi-vnfm

interface between VIM and VNFM

### ve-vnfm-em

interface between Element Manager and VNFM

### ve-vnfm-vnf

interface between VNF and VNFM