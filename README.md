# Satellogic Challenge

## Summary

- The challenge was implemented in Go.
- Observability is resolved with Prometheus for metrics, Promtail+Loki for logs, and Grafana for visualization.
- To ease testing/deployment a docker-compose is provided with all the required services and configurations.


## Running

To run the system, docker-compose must be installed. Then just run the following command in the repo root:

```bash
docker-compose up
```

## Usage

The exposed services that might be usefull are:

- An http service listening on port 8080 to add, list and execute tasks
- A Grafana instance listening on port 3000 configured with Loki and Prometheus as data sources

Down below there are instructions on how to perform some common oeprations.


### Add tasks
To add tasks make a POST request to `/tasks` with the list of tasks. Using cURL:

```bash
curl -X POST localhost:8080/tasks -d'[
    {
        "name": "capture for client 1098",
        "resources": ["camera", "disk", "proc"],
        "profit": 9.2
    },
    {
        "name": "clean satellite disk",
        "resources": ["disk"],
        "profit": 0.4
    },
    {
        "name": "upgrade to v2.1",
        "resources": ["proc"],
        "profit": 2.9
    }
]'
```

### List all loaded tasks
To list all loaded tasks make a GET request to `/tasks`. Using cURL:

```bash
curl localhost:8080/tasks
```

### Execute tasks
Execute tasks will get the list of compatible tasks that optimizes the profit and remove them from the list of pending tasks. To execute, make a POST request to `/tasks/execution`. Using cURL:
```bash
curl -X POST localhost:8080/tasks/execution
```

### View metrics and logs
Open `localhost:3000` on a browser to access the Grafana interface. Credentials are `admin/grafana` (hardcoded in the docker-compose).

A sample dashboard is provided in the repo (`o11y/grafana_dashboard.json`). It can be imported from `localhost:3000/dashboards`. It shows some metrics and logs as an example.

## Project structure

- **task_optimizer:** contains the Go code that implements the task profit optimization.
  - **internal**
      - **model:** models used by the service, right now only Task model
      - **ds:** data structures and algorithms required to solve the problem
      - **dto:** DTOs used to communicate with the service (provides abstraction between presentation/service layers)
      - **service:** implements the required methods to interact with the system (add tasks, list tasks, execute tasks)
      - **controller:** http controllers for each service method
      - **metrics:** metrics definitions for each component (allows centralization of service metrics)
      - **handler:** http handler middleware that adds logging to requests
  - **cmd:** contains the service entrypoint
- **o11y:** contains configuration files for observability components

## Algoritms and Data Structures

To solve the problem of choosing the subset of tasks that optimizes the profit, a graph of compatibility between tasks is used. In that graph, each vertex is a task and there exists an edge between tasks that doesn't share resources (tasks that are compatible). Also each vertex is weighted with the profit of the task.

Under that graph, a valid node (task) subset is that one that contains an edge between each pair of nodes. Such subgraph is called a Clique.

Modelling the problem that way, means that, to find the subset of tasks that maximizes the profit, is the same as finding the clique with maximum weight. The Bron-Kerbosch algorithm for listing all the maximal cliques was used. The provided implementation has a minor modification to output only the clique with maximum weight.

## Testing

Unit tests where added that covers 100% of the code for data structures and algorithms. Due to time constraints, it was decided to only test that part of the code.

## Possible improvements
- Use a database to store tasks. For simplicity tasks are stored in memory (they get lost once the service is restarted).
- Give more thought on which metrics are usefull to better understand the usage of the system and where to improve.
- There is an alternative Bron-Kerbosch algorithm with pivoting that reduces the search space that can be used. That might not be necessary, again, metrics can give that information.
- Improve test coverage.
- Improve logging.
- Think about which alerts to add to monitor the correct execution of the services.
- Think about alternate task models:
  - Add notion of duration estimation to each task. That allows to give priority to short-duration/medium-profit tasks instead of long-duration/high-profit tasks, for example. Maybe giving a higher profit/time.
