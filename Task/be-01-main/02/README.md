# BE - Appointment 

## Introduction
Fita is one of the company that moving in Health Industry. They have a great vision to bring the healthiness in Indonesia. To achieve that goals, they need to supervise users by providing consultation with coach feature. 
The requirement is simple, user can make an appointment with coach if only coach didn't have another appointment and have an availability. 

## Resources

Here is the coach availability data [data.csv](./data.csv)

## Requirements

The system should have an ability to:
- Make an appointment & book the schedule
- Validate an appointment based on coach availability
- Coach can decline an appointment request or reschedule it
- If user decline the rescheduling, then all ended (no need to provide another rescheduling)
- Unit test

## Intructions
Please write it in Golang or Node with a CI script that runs tests and produces a binary.

We respect your time, so you can choose four from 5 requirements but the Unit test is mandatory.

Write down the explanation and your flow in repo's README.md.

Please use git and make it public so it can be accessible by the reviewer.

Make sure always use your local time in order to make an appointment with the coach. Coach availability might not using your local time, please convert it.

