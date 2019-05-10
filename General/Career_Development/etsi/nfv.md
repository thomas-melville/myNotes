# nfv

Network Function Virtualization

Founded by 7 of the worlds leading telecoms network operators, in 2012.
the home of the Industry Specification Group for NFV

The purpose of the group is to build the Software Defined Network, SDN
Which will reduce capex, opex and time to market
NFV and SDN are complementary but increasing co-dependent

NFV Release 3 is under way, 2017 & 2018
Releases are in 2 year phases

## Working groups

### MANO

Management And Organization
Managing NFV
Working Group within ETSI, which was closed in Nov 14
It published the [ETSI NFV Architectural Framework](https://www.etsi.org/deliver/etsi_gs/NFV-MAN/001_099/001/01.01.01_60/gs_NFV-MAN001v010101p.pdf)

Mano is broken up into 3 parts

NFVO
VNFM
VIM

(They're detailed below)

### SOL

Is a new working group in ETSI NFV, SOLutions
Charged with the task of delivering a consolidated set of protocols and data model specs to support interoperability
5 work items: SOL1 -> 5

* SOL001:    TOSCA-based VNFD and NSD spec           IFA011 & IFA014
* SOL002:    REST APIs for Ve-VNFM reference point   IFA008
* SOL003:    REST APIs for Or-VNFM reference point   IFA007
* SOL004:    TOSCA-based VNF Package Specification   IFA011
* SOL005:    REST APIs for Os-Ma reference point     IFA013

### IFA

is another working group
Interfaces And Architecture
which creates functional descriptions / specs of interfaces
They have one document per reference point

* IFA005:   NFVO - VIM      Or-Vi
* IFA006:   VNFM - VIM      Vi-Vnfm
* IFA007:   NFVO - VNFM     Or-Vnfm
* IFA008:   VNFM - VNF/EM   Ve-Vnfm
* IFA009:   Architectural options for MANO components (NFVO, VNFM, VIM)
* IFA010:   Functional Requirements
* IFA011:   VNF Package & VNFD
* IFA012:   NFVO - OSS/BSS  Os-Ma-nfvo
* IFA013:   NFVO - OSS/BSS  Os-Ma-nfvo
* IFA014:   NS Descriptor

VIM n/b interfaces
...

## MANO Arch Fwk

### Basic Concepts

Relationship between VNF and E2E NS
Connectivity model between VNFs in a NS is called the VNF Forwarding Graph
VNFs can connect to PNFs
NS will have external connection points
NSs can be concatenated

Management of NFV Components
There are multiple layers and within each layer there are different types of management
On the one side there is the orchestration mgmt of each layer
and on the other there is the functional mgmt, the mgmt of the running app
The orch mgmt of each layer is grouped into NFV Management & Orchestration (MANO)

* Network Services Mgmt (SO)
* VNF Management (VNFM)
* Virtualized Resource Management (ECCD / ECM)

### Fwk

OSS/BSS     NFVO    Network Service Mgmt

EM
            VNFM    VNF Mgmt
VNF

NFVI        VIM     Virtualized Resource Mgmt

### Information Model

3 views

* application view
* * what the managed objects are
* * info model not defined by ETSI
* * they are defined by network
* logical view
* deployment view

VNF Forarding Graph and NW forwarding Path on top of Network Service

**CP**      Connection Point

* VNFs defines one or more CPs external to it
* NS defines one or more CPs external to it

**VL**      Virtual Link

* Connects two or more CPs

**VNFFG**   VNF Forwarding Graph

* The CPs and VLs define a Forwarding Graph within a NS service

**NFP**     Network Forwarding Path

* A FG can be contain multiple NFP
* If a VL has more than two connections this will enable this possibility

## Network Service

NS
A bunch of VNFs with connective and service function chaining
A network slice possibly
Network Services can be chained together to create an E2E NS

vEPC is an example of a Network Service as it is multiple VNFs

## NFVO

Network Function Virtualization Orchestrator
Manages the lifecycle of Network Services
Communicates with OSS/BSS

ECM would be an NFVO, correct. or SO???

**In a video from ETSI the NFVO is doing the onboarding of the package.**
After talking to Damian that is correct, but not all deployment scenarios will have an NFVO
We still need to adhere to the or-vnfm reference point for onboarding

## VIM

Virtualized Infrastructure Manager
Manages the NFVI resources: compute, network and storage
The clouds itself, example openstack / kubernetes, ECCD
Communicates with the VNFM

## NFVI

Network Function Virtualization Infrastructure
compute, network and storage

## VNFM

Virtual Network Function Manager
Manages the lifecycle of VNFs
When a VNF LCM request comes in the VNFM must request approval from the NFVO

## Reference Points

**ETSI plan to provide an on-line repository with a Swagger representation of the APIs!!!!**

Use TASK resources to expose complex operations

Interfaces between different "actors" in the NFV space.

Actors: VNF, VNFM, VIM, NFVO

HTTP protocol
Each reference point defines multiple interfaces between the components which the reference point links.
in both directions with HTTP Methods, query parameters, body and headers specified

### URI structure

{apiRoot}/{apiName}/{apiVersion}/

the body will be in JSON format, only!

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
These links will be in a \_links object at the same level as the omitted attribute

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

http://www.etsi.org/deliver/etsi_gs/NFV-SOL/001_099/003/02.04.01_60/gs_NFV-SOL003v020401p.pdf

Reference point which contains the interfaces between the VNFM (Manager) and the NFVO (Orchestrator)

**Interfaces.**

Component on the left produces the interface, component on the right consumes the interface

* VNF Lifecycle Management Interface. VNFM -> NFVO
* VNF Performance Management Interface. VNFM -> NFVO
* VNF Fault Management Interface. VNFM -> NFVO
* VNF Indicator Interface. VNFM -> NFVO
* * VNF indicates to the VNFM that it is running out of resources, what action to take is in the VNF Descriptor
* VNF Lifecycle Operation Granting Interface. NFVO -> VNFM
* VNF Package Management Interface. NFVO -> VNFM
* Virtualized Resources Quota Available Notification Interface. NVFO -> VNFM

### or-vi

interface between the VIM and NFVO

### vi-vnfm

interface between VIM and VNFM

Seprate APIs for compute, network and storage

* Resource Mgmt
* Info Mgmt
* change Notification
* Fault Mgmt
* Performance Mgmt
* Reservation mgmt
* Reservation change notification
* Capacity Mgmt
* Quto Mgmt
* NFP Mgmt
* Software Image Mgmt

### ve-vnfm-em

interface between Element Manager and VNFM

### ve-vnfm-vnf

interface between VNF and VNFM
