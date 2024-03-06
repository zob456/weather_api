# Weather API

### Notes
I used a `POST` request go get the weather but a `GET` request could be used & the required arguments could be passed via the URL

There are some other potential optimizations that can be done but I was pressed for time

I wrote some custom responses, loggers, etc. which are overkill for a project like this but I wrote them as a quick show of ways to create cleaner functions for other engineers to use as I have done in my current & previous roles

#### Instructions to start the application
**Prerequisite**
- [Docker desktop](https://www.docker.com/products/docker-desktop/)
- [Go v1.21](https://go.dev/doc/install) or higher to run the test command. (I use the `slices.Contains()` method which requires at least Go v1.21 I think (maybe v1.20 included it in the standard library but I'm not positive on what version introduced what outside of generics in v1.18))

**Start the API**
1. Create a `compose.yaml` file with values from the `compose.example.yaml`
2. Update the `WEATHER_API_KEY` environment value in your newly created `compose.yaml` file with your API key
3. run  `make api`
   a. This starts the API in the docker container

**Stop the application**
1. `ctrl + c` & run `make down` (or `make v-down` to remove the volumes)

**Run tests**
1. `make test {YOUR_API_KEY}`
   a. **IMPORTANT** you must pass your API key in the above test command or the tests will fail
   b. This is because I opted to use the live API rather than mock a response for time sake
