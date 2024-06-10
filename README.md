# Real-Time Notification System


## Overview

This is a real-time notification system built using the Gin web framework in Go. It allows administrators to send notifications to users. Users can log in to see notifications, and if there are no new notifications, a long-polling technique is used to receive notifications as soon as they arrive.

## Features

Admin can send notifications to users.
Users can log in to view their notifications.
Long-polling is used to provide real-time updates to users.
Secure authentication and authorization.

## Prerequisites

Go 1.16 or later
Gin web framework

## Installation

1.Install dependencies
  - go mod tidy
2.Start the server
  - go run main.go

## Usage

**1.Admin:**
  - Login to the system
  - Send notification through admin panel

**2.User:**
  - Log in to the system to view notifications.
  - Long-polling endpoint to receive real-time notifications:

## API Endpoints

1.Admin Endpoints:
**- POST /admin/login**
**- POST /admin/sendnotifications**

2.User Endpoints:
**- POST /user/login**
**- GET /user/notifications**

## Project Structure

### go

notification-system/
├── main.go
├── admin/
│   └── handler.go
│   └── routes.go
├── app/
│   └── notification.go
│   └── storage.go
│   └── signal.go
│   └── bolbol.go
├── assets/
│   └── css
│   └── javascript
├── middleware/
│   └── middlewares.go
├── templates/
│   └── html
└── user/
    └── handler.go
    └── routes.go
    
### Contributing

Contributions are welcome! Please create an issue or submit a pull request.


 
