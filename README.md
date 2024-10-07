# TutorLink

TutorLink is a custom-built tutoring platform that allows students to easily schedule tutoring sessions and make secure online payments. The platform provides a streamlined, user-friendly interface where users can view tutor availability in real-time, book sessions, and receive automated confirmations. Integrated with Stripe for secure payments and Calendly for scheduling, TutorLink ensures a seamless experience for both tutors and students.

## Table of Contents

- [Features](#features)
- [Tech Stack](#tech-stack)
- [Usage](#usage)
- [License](#license)
- [Contact](#contact)

## Features

- **Online Scheduling**: Real-time session booking integrated with Calendly for syncing tutor availability and student appointments.
- **Secure Payment Processing**: Integrated Stripe API for secure online payments.
- **User Management**: Account creation and profile management for students and tutors, allowing personalized user experiences.
- **Responsive Design**: Optimized for mobile, tablet, and desktop, ensuring users can access the platform from any device.
- **Email Notifications**: Automated email confirmations and reminders for scheduled sessions and payments.
- **Scalable Deployment**: Built with Go and React (Vite), ensuring fast, high-performance delivery and ease of scaling.
- **Atomic Session Management**: Handle session booking, cancellations, and payments seamlessly with error handling for no-shows and reschedules.

## Tech Stack

### Front-End

- **React (Vite)**: For building a responsive and interactive user interface, bundled and served with Vite for fast development and optimized builds.
- **JavaScript**: Core language for building React components and logic.
- **React Router**: For handling routing and navigation within the app.
- **Vite**: A blazing-fast build tool that serves the React app in development and production.
- **Axios**: For handling HTTP requests to the backend API.
- **CSS Modules** or **Tailwind CSS**: For styling the app with scoped or utility-first CSS.

### Back-End

- **Go (Golang)**: Backend built with Go for high performance, concurrency handling, and scalability.
- **Gin**: A fast, lightweight HTTP web framework for building the API.
- **PostgreSQL**: Relational database for managing user accounts, session data, and payment records.

### API

- **REST API**: A RESTful API built with Go and Gin for handling client-server communication.
- **Stripe API**: For handling secure payments, refunds, and session payment policies (no-shows, cancellations, etc.).
- **Calendly API**: For managing tutor availability and scheduling student appointments.

### Scheduling and Payment Integration

- **Calendly**: For scheduling and managing tutor-student sessions.
- **Stripe**: For secure online payment processing.

### Authentication

- **JWT (JSON Web Tokens)**: For secure, stateless user authentication.

### CI/CD and Deployment

- **GitHub Actions**: For automated testing, building, and deployment.
- **Docker**: Containerization of both backend and frontend for consistent deployment environments.
- **Google Cloud Run**: For deploying the Go backend, providing scalable and serverless deployment.
- **Vercel** or **Netlify**: For deploying the React frontend for fast, globally distributed content delivery.

### Testing

- **Go Testing**: For testing backend business logic and API endpoints.
- **Jest**: For unit and integration testing of the frontend.
- **React Testing Library**: For testing React components and UI behavior.
- **Supertest**: For API testing from the frontend during development.

## Usage

### Running Locally

1. **Clone the repository**:

   ```bash
   git clone https://github.com/your-repo/tutorlink
   cd tutorlink
   ```

2. **Run the backend**:

   - Navigate to the backend directory:
     ```bash
     cd backend
     ```
   - Ensure Go is installed, then run:
     ```bash
     go run cmd/api/main.go
     ```

3. **Run the frontend**:

   - Navigate to the frontend directory:
     ```bash
     cd frontend
     ```
   - Install dependencies and start the Vite development server:
     ```bash
     npm install
     npm run dev
     ```

4. **Access the app**:
   - Frontend: `http://localhost:3000`
   - Backend (API): `http://localhost:8080/api`

### Docker (Optional)

You can also run both the backend and frontend using Docker:

1. **Build and run using `docker-compose`**:
   ```bash
   docker-compose up --build
   ```

This will start both the frontend and backend in containers.

## License

TutorLink is licensed under the MIT License. See the [LICENSE](LICENSE) file for more details.

## Contact

Alexis San Javier - [ucfknight2017@gmail.com](mailto:ucfknight2017@gmail.com)

- **Website**: [alexissj.net](https://www.alexissj.net)
- **LinkedIn**: [linkedin.com/in/alexissj](https://linkedin.com/in/alexissj)
