#!/usr/bin/env python
# coding:utf-8
# Filename:mongodb.py

from pymongo import MongoClient,ASCENDING,DESCENDING
import datetime

# connection with mongoclient
client=MongoClient()

# getting a database
db=client.test

# getting a collection
collection=db.posts

# documents
post={"author":"Mike",
      "test":"My first blog post!",
      "tags":["mongodb","python","pymongo"],
      "date":datetime.datetime.utcnow()
}

# inserting a document
post_id=collection.insert(post)
print 'posts id is:',post_id
print 'collection_names is:',db.collection_names()

# getting a single document
doc=db.posts.find_one()
print doc

#query by objectId
print 'query is:', db.posts.find_one({"_id":post_id})

# querying for more than one doc
for post in db.posts.find():
    print post

# counting
print 'total count is:',db.posts.count()

# range queries
d=datetime.datetime(2014,8,9,12)
for post in db.posts.find({"date":{"$gt":d}}).sort("author"):
    print 'gt is:',post

# Indexing
# before indexing
print db.posts.find({"date":{"$gt":d}}).sort("author").explain()["cursor"]
print db.posts.find({"date":{"$gt":d}}).sort("author").explain()["nscanned"]
# after indexing
db.posts.create_index([("date",DESCENDING),("author",ASCENDING)])
print db.posts.find({"date":{"$gt":d}}).sort("author").explain()["cursor"]
print db.posts.find({"date":{"$gt":d}}).sort("author").explain()["nscanned"]

# remove all indexes
db.posts.drop_indexes()
