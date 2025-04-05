# Agenda Manager

Manage agenda, person availability and integration with google calendar

## General terms

- Promoter: the person who promotes an event or its the owner of an availability agenda.
- Attendee: the person who will participate in the meeting and/or schedule a meeting using the availabilty agenda.

## Technical requirements 

- API RESTful
- Authentication (using google)
- Authorization
- Postgres
- Google calendar integration

## Requirements

### Promoter 

- Availability Agenda
  - Create an availability agenda
    - Timezone 
    - Weekdays (that the person will be available)
    - Available time per weekday
    - Timeslot limit
    - Limit event amount per person
    - Which frequency the limit resets ?
      - Lifetime
      - Weekly
      - Monthly
      - Biweekly
      - Anually
    - Google calendar integration
      - If the time is blocked/busy on my google agenda it should not appear on my availability agenda.
     
### Attendee

- Schedule an event
  - First name
  - Last name
  - Email
  - Description
  - Using an availabilty agenda from a Promoter I need to be able to schedule an event.
  - If the attendee does not any other appointments left available it will return an error.
 
## Implementation Steps

1. Define base models.
2. Implement the data access layer. (Use a repository pattern)
3. Implement the logic of data comparison.
4. Implement the logic to create agenda.
5. Implement the logic to create an event using an existing agenda.
6. Implement logic to retrieve user info from google api.
7. Implement api with access control to the endpoints
8. Document api using Swagger

## Stack to use

- Golang
- Gin Framework
- Postgres
- Swagger
- Google API

