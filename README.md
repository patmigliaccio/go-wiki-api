# untitled

## Deployment

### Local Deployment

Docker Build

```bash
docker build -t untitled .
```

Docker Run

```bash
docker run --publish 7200:80 --name untitled --rm -d untitled
```

Accessible at `http://localhost:7200`

### Production Deployment

Prerequisites

```bash
yarn global add now
```

Now Deployment

```bash
now
```