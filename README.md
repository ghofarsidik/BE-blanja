<div align="center">
  <a href="https://github.com/dimassagngsptr/react-client-blanja.git">
      <img src="https://github.com/dimassagngsptr/blanja-kelompok-1/blob/master/apps/web/src/assets/logo/Group%201158.png" width="300"/>
  </a>

  <h1 align="center">Blanja Innovation</h1>

  <p align="center">
     Blanja Implementation
    <br />
    <br />
    <a href="https://innovation-blanja.vercel.app/" target="_blank">View Website Demo</a>
    ·
     <a href="https://github.com/dimassagngsptr/golang-server-blanja.git" target="_blank">View Back-End Repo</a>
  </p>
</div>

- [Table of Contents](#table-of-contents)
- [About The Project](#about-the-project)
  - [Built With](#built-with)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
- [Contributing](#contributing)
- [Contact](#contact)
- [Documentation](#documentation)
- [Acknowledgements](#acknowledgements)




## About The Project

Blanja is a sales application specializing in clothing items, including jackets, pants, t-shirts, and shirts. The application is built using React.js for the client side and Golang for developing the RESTful API. With the assistance of the Gofiber framework, Blanja offers functionalities such as search, authentication, authorization, add-to-cart, customer address addition, and checkout features utilizing the payment gateway from Midtrans.

### Built With

[Golang](https://go.dev/)

[Gofiber](https://gofiber.io/)

[GoORM](https://gorm.io/)

[Postgre Sql](https://www.postgresql.org/)

[Cloudinary](https://cloudinary.com/)

[Midtrans](https://midtrans.com/)

[JSON Web Token](https://jwt.io/)

[Bluemonday](https://github.com/microcosm-cc/bluemonday)

## Getting Started

To get a local copy up and running follow these simple steps.

### Prerequisites

Ensure you have the following installed on your local machine:

- Go
- PostgreSQL

### Installation

1. Clone Repo

   ```sh
     git clone https://github.com/dimassagngsptr/golang-server-blanja.git
   ```

2. Go to workspace

   ```sh
      cd golang-server-blanja
   ```

    Open VsCode

    ```sh
      code .
    ```

3. Running the project

   ```sh
     air
   ```

The server will start on port 3000 by default. You can use Postman to interact with the endpoints in [Documentation](#documentation).

## Usage

To use this project, follow the instructions below for understanding the project structure and how to use the provided API documentation.

### Project Structure

```
golang-server-blanja/
 .air.toml
│   .env
│   .gitignore
│   go.mod
│   go.sum
│   main.go
│
├───src
│   ├───configs
│   │       db.go
│   │
│   ├───controllers
│   │       addressController.go
│   │       authController.go
│   │       cartController.go
│   │       categoryController.go
│   │       productColorController.go
│   │       productController.go
│   │       productSizeController.go
│   │       storeController.go
│   │       transactionController.go
│   │       userController.go
│   │
│   ├───helpers
│   │       countData.go
│   │       generateJwt.go
│   │       generateRandomNumber.go
│   │       migration.go
│   │       transformFilePath.go
│   │       uploadFile.go
│   │       validationPassword.go
│   │       validator.go
│   │       xssClean.go
│   │
│   ├───middlewares
│   │       jwt.go
│   │
│   ├───models
│   │       Address.go
│   │       Cart.go
│   │       Category.go
│   │       Product.go
│   │       ProductColor.go
│   │       ProductImage.go
│   │       ProductSize.go
│   │       Store.go
│   │       Transaction.go
│   │       User.go
│   │
│   ├───routes
│   │       main.go
│   │
│   ├───services
│   │       cloudinary.go
│   │       midtrans.go
│   │       smtp.go
│   │
│   └───templates
│           email_template.html
│
└───tmp
        main.exe
```

### Documentation

Access the API documentation for the **Blanja** project. Use this documentation to test endpoints and understand the structure and functionality of the available APIs in this project.

[![Blanja Innovation API Postman Documentation](https://run.pstmn.io/button.svg)](https://documenter.getpostman.com/view/29288801/2sA3duECJC)


### Project Related

- [`Front-End Project Repository Link`](https://github.com/dimassagngsptr/react-client-blanja.git)

- [`Front-End Demonstration Link`](https://innovation-blanja.vercel.app/)

- [`Back-End Demonstration Link`](https://golang-server-blanja-production.up.railway.app/)


## Contributing

Contributions are what make the open source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## Contact

If you have any questions or inquiries regarding this project, feel free to contact:

email : dimasageng58@gmail.com 

linkedin : https://www.linkedin.com/in/dimasagengsaputro/

## Acknowledgements

Feel free to check it out:

- [Img Shields](https://shields.io)
- [Choose an Open Source License](https://choosealicense.com/)
- [GitHub Pages](https://pages.github.com/)
