# Infrastructure layer

The library's infrastructure layer provides a cohesive mechanism for storing events as sequences of indexed event log entries.

The entire mechanism is encapsulated by the EventStore. The event store uses a sequenced item mapper and a record manager.

The sequenced item mapper converts objects such as domain events to sequenced items, and the record manager writes sequenced items to database records. The sequenced item mapper and the record manager operate by reflection off a common sequenced item type.

