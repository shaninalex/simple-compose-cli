# Simple console cli

Simple managing docker compose projects.

### Usage

```bash
# list
$ ./bin/compose-cli --list
project-name-1      ( /<absulute path>/docker-compose.yml )
project-name-2      ( /<absulute path>/docker-compose.dev.yml )
project-name-3      ( /<absulute path>/docker-compose.test.yml )
project-name-4      ( /<absulute path>/docker-compose.yml )

# start
$ ./bin/compose-cli --name=project-name-1 start
/usr/local/bin/docker compose --file /<absulute path>/docker-compose.dev.yml --project-name project-name-1 stop
Done

# stop
$ ./bin/compose-cli --name=project-name-1 stop
/usr/local/bin/docker compose --file /<absulute path>/docker-compose.dev.yml --project-name project-name-1 start
Done

```