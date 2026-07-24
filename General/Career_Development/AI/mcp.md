# MCP

Open standard created by Anthropic.
Model Context Procotol.
A standard for connecting large language models to relevant data and tools.

Up to now, developers have had to provide meticulous details of context in prompts.
MCP allows LLMs and agents to interact with other dataset to get more information for a particular task.

TO integrate an MCP add an mcp.json file to your project to start an mcp server

A communication layer to provide agent with context and tools without you having to write a lot of code.
An interface to an outside service.

Main elements of MCP
    MCP Client
    MCP Server

Server contains some internal components.
    Tools
    Prompts
    Resources

MCP Client to Server communicate on Standard IO, usually they are running on the same machine.
Can also communicate many other ways, HTTP, websockets, ...

Then communicate by exchanging messages, which the schema of is defined in the spec.

There's an official mcp python module which is an SDK for creating MCPs an tools.

## server

mcp sdk has python APIs to create a MCP quickly.
import mcp

FastMCP class to create an mcp instance.

mcp.tool decorator to create a tool in the MCP

Pydantic is also used when developing MCPs to give more information to tool Schemas

mcp.resource decorator, allow for auto complete on resources which the client asks the server for.
    direct resource - static URI resource in the decorator
    tempalted resource - URI contains one or more params.

mcp.prompt(name, description) decorator
    Define a set of User and assistance messages that can be used by clients.
    prompts can be high quality, well tested and relevant to the overall purpose of the MCP
    function implementation returns a prompt with any inputs filled in
    turns up in the client as a / command.


## tools

name, description, arguments and return type.
Pydantic arguments help the MCP SDK define json schemas which are then shared with the MCP Client.

## inspector

Handy tool to run and test your MCP locally without having to set up the full flow.

## client

async module is used when interacting with mcp server through the session

MCP Client wraps a session to the MCP Server, Session class is part of the MCP SDK.
MCP Client is used to get list of tools and run tools from MCP Server.
session.list_tools(...)
session.call_tool(...)

