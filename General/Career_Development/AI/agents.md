# agents

The shift is to writing agents to use LLMs which then provide the user with the information they want.
The agent constructs the prompt and context.
They seem to be mostly written in Python

Receive, decide, act

## lifecycle

different to a normal app.
AI agents evolve with context overtime!

Agents need to be observed one deployed

prompts will need to be tuned, address hallucinations

agents will eventually need to be retired
    no longer used, BL changes, model outdated, 

agent manifest:
    prompt versions
    model used, including versions
    tool config
    environment settings
in yaml, can have a schema for it.
    Doesn't sound like it's standard

## dependencies

APIs which the agent talks to
Vector DBs
LLM Models
authentication providers
retrieval indexes

## best practices


## multi agent system

Agents are good at doing a single task.
When you have a more complex workflow, you can string multiple agents together
Workflow can be linear or like a graph
LangGraph was created to manage the entire lifecycle and flow of multi agent systems.

### LangGraph

Code libraries, defines patterns for agents.
Helps agents remember and communicate with other agents using state management 
Test and debug AI agents
Deployment and scaling
    packages as docker container
tracing

For AI agents state is the info that's available to process.
agent remembers what you've told it for context.
LangGraph has means to stoe this, so agent knows what other agents have done.

Think of work flow as a state machine.
Defined actions and transitions.

LangGraph defines workflows using Graphs
Nodes represent actions, edges where to go next
sequential, branching (AI function makes a decision), cyclic, multi agent

Testing in LangGraph
    Studio, generates diagram of workflow
        live debugging
    Traces, capture data through workflow
    LangSmith, setup experiments to compare outputs before and after agent changes
    LLM MEtrics, 

## MCP

Model Context Protocol, how different components in an AI ecosystem can communicate and share context.
This is getting into the area of multi-agent, where to complete the task multiple agents are employed.


## sub agents

an agent can call sub agents to execute tasks.
sub agents have access to same tools as parent that calls them.
Two types of sub agents
1. map reduce
    break work into different chunks
    process tasks simualtaneously
2. task dependency graph
    sequential
kiro will know what way to organize work of sub agents
they can be combined in a single execution depending on the work

agent configuration, add "subagent" to tools and allowedTools arrays
