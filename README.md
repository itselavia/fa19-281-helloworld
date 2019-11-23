# Team Hackathon Project: University Portal

## Team Members

| Name | Microservice |
| ------ | ----------- |
| Akhilesh Anand    | Profile Service |
| Akshay Elavia  | Admin Service |
| Arivoli AE    | Enrollment Service |
| Nirbhay Kekre    | Course Service |
| Samarth Khatwani    | Grading Service |

## Summary
+ **University Portal** is a cloud based Software-as-a-Service application similar to MySJSU for managing academic activities of a University.
+ Students can sign up and perform the following actions:
  - Search for Courses
  - View Availability of Seats
  - Enroll into courses
  - Pay Fees for the term
  - View Grades
  - View Announcements
  - Dismiss Announcements
+ Instructors can sign up and perform the following actions:
  - Create a Course
  - Submit Grades for a particular course
  - Broadcast Announcements
  - View Analytic info such as:
    - Most Searched Courses
    - Best Performing Courses    
## High-level Architecture Diagram
![Alt text](https://github.com/nguyensjsu/fa19-281-helloworld/blob/master/mysjsu%20(1).png)

## Summary Key Features
* **Kubernetes deployment:** [more details](https://github.com/nguyensjsu/fa19-281-helloworld/blob/master/extra-credit-kuber-award.md)
* **Countinous deployment:** [more details](https://github.com/nguyensjsu/fa19-281-helloworld/blob/master/extra-credit-DevOps-award.md)
* **Event Sourcing via Kafka:** Kafka was leveraged to communicate asynchronously over the queue between microservices [more details](https://github.com/nguyensjsu/fa19-281-helloworld/blob/master/extra-credit-kuber-award.md)
* AFK cube dimensions:
  * **X axis:** horizontal scalling with multi docker instance deployment of backend.
  * **Y axis:** Functional split into microservices.
  * **Z axis:** Splitting data into shards with replicasets.
