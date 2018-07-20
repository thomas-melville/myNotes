# exploratory testing

guided manual testing, with tours

all systems do 4 things:

1. accept input
2. produce output
3. store data
4. perform computations

positive V negative testing

## input data

what is illegal input? as defined by the developer

### error handlers

1. input filters
  * limit what the user can actually enter, ex: drop down boxes
  * is what is constituted as legal / illegal correct?
  * Can it be bypassed?
2. input checks
  * if / then / else or switch / case structures
  * pay attention to the information given to the user when they enter illegal input!
3. exceptions
  * handlers are more generic than input checks, will handler other kinds of failures

### normal / special input

a normal input has no special formatting or meaning and is just used by the app.
a special input is something random that the app does not expect, it's more than illegal input

* think of Ctrl, Alt, Esc, ...

### default / user supplied data

null / no data may result in default case being executed!
Default values may mean there was no thought given to entering no data!

### using outputs to guide input selection

think of the outcome/output you want and put in the data you think will give you this output
input the same data several times to see how the system works when it is new and then dirty (has state/data from the previous execution)

## state

if an app stores state then entering the same data twice can have 2 different results:

1. When the app is not in the state it ends up in.
2. When the app is already in the state that it would end up in

the state space of an app is huge!
It is the accumulation of all possible values of each persistent object/value/...

we need to test the scope of data.

Persisted for:

* the session
* the uptime of the app
* ever, in a database

combine related input data and state to thoroughly test functionality

## the metaphor

A tourist in a city for the first time, how to make the most of your time!
Aimlessly wander, book tours, get a guidebook, ...

### districts

cities are divided into districts in Guidebooks

* Business
  * in code, where the business gets done
  * this is the start up, functionality and shut down of the app
* Historical
  * for sw, legacy code and history of buggy functions & features
  * legacy code is often poorly understood and assumptions are made about it
* Tourist
  * in sw, novice users are attracted to features and functions that experienced users have no more need for
* Entertainment
  * supportive features which prop up the main functionality
* Hotel
  * when the software is at rest
* Seedy
  * do things not recommended
  * find vulnerabilities

### tours

#### business district

These tours focus the tester on the important features and guide testers through the paths of the software where these features are used.

##### The Guidebook Tour

Attractions in a guidebook are the most popular and tourists are given recommendations for how to enjoy them.
They are the most popular so they must be clean and safe.

Start with the user manual for this tour, sometimes this tour is called the F1 tour (the short cut key for the help menu)
Follow the user manuals advice, never deviate from it's lead. Follow it to the letter

##### The Money Tour

Where's the money? What is making customers buy your software???

Look at all sales material and demos. Try and execute them all

##### The Landmark Tour

Orient yourself by using landmarks which are in the general direction of where you want to go

Choose a set of landmarks in the app and hop through them in different orders.

##### The Intellectual Tour

Ask a tour guide hard questions to see if he knows what he is talking about

Ask the software hard questions, make it work as hard as possible.

* which features will stretch it to its limits
* what input & data will cause it to perform the most processing
* what inputs might fool its error-checking routines
* which input and internal data will stress it's capability to produce any specific output

##### The FedEx Tour

FedEx is an icon in the package delivery world. pick up here, deliver there.

Think of data moving through the software.

* It is input
* stored internally
* manipulated, modified and used in computations
* delivered as output

identify inputs that are stored and follow them through the system

##### The After-Hours Tour

After the business day is done, everyone goes home. This time is very busy so tourists tend to stay away

Even when the software isn't being used it is at work.

* maintenance tasks
* archiving data
* backup

remember to test these things, possibly by forcing them to occur

##### The Garbage Collectors Tour

Garbage collectors cover an entire neighborhood methodically, house by house.
But because they are in a hurry they don't stay anywhere long.

a methodical spot check in software.
screen by screen / dialog by dialog the shortest route possible
don't test in details
choose a goal of what you want to check and visit all of those in the shortest route possible

#### historical district

These tours focus the tester on legacy code, early features and bug fixes. test legacy functionality & verify bug fixes

##### The Bad Neighbourhood Tour

every city has it's bad areas that tourists are advised to stay away from.
Exploratory testers should do the opposite.
At first you can't know where the buggy areas are so keep a record of where bugs are found.
And the feature/area with the highest bug count is the bad neighbourhood

##### The Museum Tour

Antiquities are a favourite of tourists.
Softwares antiquities are legacy code.

Use version control to find out the least touched area of code.
older code with recent modifications.

##### The Prior Version Tour

run all scenarios and test cases that applied to the prior version.
sounds like a regression test?
func reimplemented / removed you need to test for new way of doing things

#### entertainment district

These tours focus the tester on supporting features, nice to haves that aren't used by all users. They're not the core features that are the selling point of the product

##### The Supporting Actor Tour

On a guided tour there are things around what the tour guide is pointing to that interest people.
When a sales person demos a product, customers are also looking at features which are nearby the spotlight

Test features which share the screen with the features we expect the users to use most.
Their proximity to the main event increases their visibility