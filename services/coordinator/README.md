# Concierge

- when items such as experiences, memories, gems, etc. are created, the concierge will index these items.
- when the client calls for a curation the concierge will create one on the fly based on the params sent by the client as well as the guest's persona.
- The curation will contian a list of items related to the query
- the client will then display the items appropriately
- on subsequent pages get the curation data from the cache and parse it for the experience that should be displayed.
-

## Startup

The concierge depends on other services at startup

## Start Operations

- load env
- start postgres db (docker)
- start redis-stack-service
- start service
- start journeys

## TODOs

- install <https://github.com/laixintao/flameshow>

params.connStr = "host=localhost user=tony password=password dbname=concierge port=5432 sslmode=disable"

## Resource creation

guest creates a journy: only guests can create ??

- for self
- for other guest (there's who did it and for who they did it for)
  - parnter creates (still a guest)
    - guest calls partner and says "hey, can you create a resource for me?"
    - a guest can reach out to another guest on the platform (an expert) and
    - asks them to create a resource.. that expert getes a peice of the pie.
  - admin/help desk (still a guest)

dorathy suggested resource (not an actual resource, just a suggestion)

- there is also the idea that each guest will have their own 'child' of dorathy where it will be able to assist
  the guest with a number of tasks that are relavent to them and only them (giga-pets!!).

## Creating experiences

- Client will make a request to the concierge containing the resource and experience info in the payload.
- the concierge will return to the client a request id and the status of the request/job
- the resource service create the rules and data associated with the experience such as the required resources, pricing, and availability.
- the journey service will save the experinece info.
- the concierge, via its content module, will save the images associated with the experience/resource(s)
- in the event that the job was not completed the client will be able to use the requeest id to check the sstatus of the job.
- The client should be able to query the concierge for jobs.

### Errors

Experience name is not unique

- The concierge will return the request id along with an appropriate error message.

- the client will need to use that request id in future calls if they are in fact in the process of creating the same experience.
-
- the concierge will make a call to the operations sevice to validate and create the required resources.

- an event is saved to the store and propagated to downstream services such as Dorathy.

## Searching

- When a guest queries for a resource or
- Searching is handled by Dorathy.
- Dorathy listens for various events throughout the system and updates its own db in real time.
- when a guest searches for a resource, experience, or a memory, that request will be passed to Dorathy via the concierge.
- Dorathy will then return what's found back to the guest client.
- while dorathy is able to provide realtime data its data cannot be depended on for availability.
- dorathy may return an experience as being available (how the hell to handle this.)

## Research

https://aws.amazon.com/swf/
https://netflix.github.io/mantis/
https://cadenceworkflow.io/docs/get-started/#what-s-next
