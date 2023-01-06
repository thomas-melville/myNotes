# accessing data

## open a file and json

```python

d = open('filename', 'rt')

all_j = d.read()

all_l = json.loads(all_j)

```

rt means read text, text is default so you don't need to add it.

## remote file loading

using requests

response has method json() to read the response body as a json object

## convert python structure to json

```python

a = [] # python structure

a_j = json.dumps(a)

```

## use json to encode a class

A class needs getters and setters for fields

```python

class Item:

class ItemEncoder(json.JSONEncoder):
  def default(self, obj): # pass in the object to be encoded
    if isinstance(obj, Item): # if it's an instance of our class then return whatever we wish to encide
      return onj.__dict__
    else: # otherwise use the default encoder, or you could/should raise an aexception
      return json.JSONEncoder.default(self, obj)

z_j = json.dumps(z, cls=ItemEncoder)

```

**or just serialize the dict of any object!!!**

json.dumps(z.__dict__)

But.... encode allows you to be specific about what you encode.
  don't encode certain fields
  change type
  ...

## database

You should wrap database interactions in try/except blocks

```python

import sqlite3 # built in to Python!!!

conn = sqlite3.connect('my_db') # connect/create database

curs = conn.cursor() # ponting to a point in the database/result set.

st = '''
    CREATE TABLE zoo(
      ...
    )
'''

curs.execute(st)
conn.commit()
conn.close()

st = '''
    INSERT INTO zoo VALUES (?, ?, ?)
'''

curs.execute(st, ("string", 25, 1.56))

```

cursor is used to execute statements and fetch results.

cursor.fetchone/many/all
returns a list of tuples
