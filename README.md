# Task

## Run the application

```
npm i --prefix client
docker compose up --build
```

## Project overview

This is a very simple app built using SvelteKit, with a ready-to-use Go server.
It uses an open API to download a list of anime and display them. You can access anime details and add them to favorites.
All favorites are then displayed on the right side of the screen.

## IMPORTANT

This API has a rate limiter, so sometimes there might be some errors. Just refresh and try again.

1. Copy this repo (don't fork it!) into Your github account.
2. Create a new branch on Your github and start working there.
3. After finish, create a Pull Request from Your branch to Your Main and invite me.  
   This way i will be able to see all the changes, and my repo will be clean for the next candidate :).
4. Please check that Your example can be run without any additional configuration:

- delete repo
- clone it again
- switch to Your branch
- run the command

```
docker compose up --build
```

If you will have any questions, just write to me:
`mateuszpiorowski@gmail.com`

## Tasks

Please, try to not spend more than 6 hours.
You don't need to fully finish all of them, choose how much you want to dive into each of them.

**But You must at least touch all of them. Sorry :)**

The Backend part might be the hardest, so if it will start to take too much time, just show some sql knowlegde or write one simple endpoint.

### Styling

Style this application a little bit using TailwindCss.
Add some animation using gsap (preferably), animejs, motion or svelte animation.
Add a table with pagination to the anime list.

### Functionality

On the anime details page, add a button to remove the anime from favorites.
For the "Add to favorites" form action, add validation on the server side using the Zod library.
Make it that there is no type assertion ("as") anywhere.

### Backend

Instead of using the data/favorites.ts object, use a connected Go backend to store and load favorites.
The server is connected, with a workign SQLite database and a simple API.

A little helper:

1. Add necessary columns to the favorites table ( `system_migration.go` )
2. Add INSERT and DELETE queries to the server ( `favorites_db.go` )
3. Add new service layer to the server to add / remove favorites ( `favorites_http.go` )
4. Add new endpoints to the server ( `main.go` )

To clean the database, just remove the `db.sqlite3` files inside `system` folder and restart the Docker container.

## Reference Links

- GSAP: https://gsap.com/
- Svelte: https://svelte.dev/
- SvelteKit: https://kit.svelte.dev/
- Svelte Form Actions: https://kit.svelte.dev/docs/form-actions
- Go By Example: https://gobyexample.com/
- Go Tour: https://go.dev/tour/welcome/1
