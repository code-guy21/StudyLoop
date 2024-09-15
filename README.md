# TutorLink

TutorLink is a custom-built tutoring platform that allows students to easily schedule tutoring sessions and make secure online payments. The platform provides a streamlined, user-friendly interface where users can view tutor availability in real-time, book sessions, and receive automated confirmations. Integrated with Stripe for secure payments and Google Calendar for scheduling, TutorLink ensures a seamless experience for both tutors and students.

## Table of Contents

- [Features](#features)
- [Tech Stack](#tech-stack)
- [License](#license)
- [Contact](#contact)

## Features

- **Online Scheduling**: Real-time session booking integrated with Google Calendar API to sync tutor availability and student appointments.
- **Secure Payment Processing**: Integrated Stripe API for secure online payments, including one-time and package-based payments.
- **User Management**: Account creation and profile management for both tutors and students, allowing personalized user experiences.
- **Responsive Design**: Optimized for mobile, tablet, and desktop, ensuring users can access the platform from any device.
- **Email Notifications**: Automated email confirmations and reminders for scheduled sessions and upcoming payments.
- **SEO Optimized**: Enhanced for search engines with server-side rendering (Next.js), ensuring high search rankings and better discoverability.
- **Scalable Deployment**: Hosted on AWS or Google Cloud Platform for scalable, high-performance delivery with global reach.

## Tech Stack

### Front-End

- **React**: For building a responsive and interactive user interface.
- **TypeScript**: For type safety and better developer experience.
- **Redux Toolkit**: For managing global state and handling user sessions.
- **React Router**: For seamless navigation and routing.
- **Next.js**: For server-side rendering to enhance SEO and performance.

### Back-End

- **Node.js** with **TypeScript**: For server-side scripting and type safety.
- **Express.js**: For building a RESTful API.
- **PostgreSQL** with **TypeORM**: For relational data management, including user accounts and session scheduling.

### API

- **REST API**: For handling client-server communication.

### Scheduling and Integration

- **Google Calendar API**: For syncing tutor availability and scheduling student appointments.

### Payment Integration

- **Stripe**: For secure online payment processing.

### Authentication

- **JWT (JSON Web Tokens)**: For secure, stateless user authentication.
- **OAuth 2.0** (Optional): For integrating social login options like Google or Facebook.

### CI/CD and Deployment

- **GitHub Actions**: For automated testing, building, and deployment.
- **Docker**: For containerization to ensure consistent environment setup.
- **AWS Elastic Beanstalk** (or **Google Cloud Platform**): For scalable, high-performance deployment and management.

### Testing

- **Jest**: For unit and integration testing of both backend and frontend.
- **React Testing Library**: For testing React components.
- **Supertest**: For API endpoint testing.

## Usage

1. **Sign Up**: Register as a student or tutor to start using the platform.
2. **Set Availability**: Tutors can set their availability, which is synced with Google Calendar.
3. **Book Sessions**: Students can view tutor availability in real-time and book sessions directly.
4. **Make Payments**: Securely pay for sessions using Stripe, with options for one-time or package-based payments.
5. **Receive Notifications**: Get automated email confirmations and reminders for scheduled sessions and payments.

## License

TutorLink is licensed under the MIT License. See the [LICENSE](LICENSE) file for more details.

## Contact

Alexis San Javier - [ucfknight2017@gmail.com](mailto:ucfknight2017@gmail.com)

- **Website**: [alexissj.net](https://www.alexissj.net)
- **LinkedIn**: [linkedin.com/in/alexissj](https://linkedin.com/in/alexissj)
