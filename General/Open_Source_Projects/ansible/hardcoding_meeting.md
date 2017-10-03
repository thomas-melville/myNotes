"hardcoding" meeting

Current situation
	networkElementId is duplicated across data in testware
	A simple example is NetworkExplorer, the networkElementId is in the data for what nodes to search for and add to a collection.
	A complex example is CM Import, the networkElementId is in the xml files for import.

what's wrong with that?
	If NSS change the teams allocation their testware needs updating.
	If NSS execute the testware with a smaller NRM the tests will fail.

possible solutions
	merge any and all datasource used by the team with the nodesToAdd central csv to filter down to the correct number of nodes
	pop the networkElementId's into properties and reference in data with ${} so that they are always using the correct values
	move away from networkElementId to neType, ???
	separate profiles per each netsim environment / nrm


short term solution
	separate profiles
		combined with one or more of the above solutions

long term solution
	move to tdm and neType as the identifier

	this requires:
		tdm to be released. Robbie and Ger to test in staging
		a new requirement: lock a single data record. Double check does Vlad's locking story lock individual records

new ticket: usage metrics for tdm