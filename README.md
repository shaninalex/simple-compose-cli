# Simple console cli

Simple managing docker compose projects.

### Usage

```bash
# list
$ ./bin/compose-cli --list
project-name-1      ( /<absolyte project path>/docker-compose.yml )
project-name-2      ( /<absolyte project path>/docker-compose.dev.yml )
project-name-3      ( /<absolyte project path>/docker-compose.test.yml )
project-name-4      ( /<absolyte project path>/docker-compose.yml )

# start
$ ./bin/compose-cli --name=project-name-1 start
/usr/local/bin/docker compose --file /<absolyte project path>/docker-compose.dev.yml --project-name project-name-1 start
Done

# stop
$ ./bin/compose-cli --name=project-name-1 stop
/usr/local/bin/docker compose --file /<absolyte project path>/docker-compose.dev.yml --project-name project-name-1 stop
Done

```