# nobones-api
A bare-bones API to determine if today is a [bones day](https://knowyourmeme.com/memes/bones-day-no-bones-day)

# Usage
This app provides several endpoints that can be used depending on your needs.

## JSON Response
The `GET /` endpoint returns a JSON object with a single boolean: `bones`. This will be true on "bones" days, and false on "no bones" days.

```
> https nobones.today
HTTP/1.1 200 OK
Content-Length: 15
Content-Type: application/json; charset=UTF-8
Date: Sat, 30 Oct 2021 12:44:53 GMT

{
    "bones": true
}
```

## Bones Status
If parsing JSON is too much, the `GET /bones` endpoint is for you. This will return an empty `200` response if it's a "bones day." Likewise, on "no bones" days, this will return a `404` status.

```
> https nobones.today/bones
HTTP/1.1 200 OK
Content-Length: 0
Date: Sat, 30 Oct 2021 12:52:45 GMT

```

## Kubernetes Admission Webhook
The `POST /admission` endpoint is a Kubernetes [validating admission webhook](https://kubernetes.io/docs/reference/access-authn-authz/extensible-admission-controllers/) that will only authorize requests on "bones days." This is super helpful if you wanted to easily prevent application deployments on "no bones" days.

```
> https POST nobones.today/admission @test/request.json
HTTP/1.1 200 OK
Content-Length: 201
Content-Type: application/json; charset=UTF-8
Date: Sat, 30 Oct 2021 12:55:41 GMT

{
    "apiVersion": "admission.k8s.io/v1",
    "kind": "AdmissionReview",
    "response": {
        "allowed": true,
        "status": {
            "code": 200,
            "message": "It's a bones day!",
            "metadata": {}
        },
        "uid": "705ab4f5-6393-11e8-b7cc-42010a800002"
    }
}
```

# Data Source
Using the O.G. key-value store, DNS text records, I've created a data source for this API and have made it public. I'll likely keep that record updated far longer than a reasonable person would.

```
> dig +short TXT api.nobones.today
"bones"
```

# Build/Run
This is a small Go application shipped as a container. You can build it yourself or use my prebuilt image at `quay.io/patrickeasters/nobones-api`.

```
> docker build -t nobones-api .
docker run -p 3000:3000 -it nobones-api
[+] Building 0.3s (10/10) FINISHED
...

> docker run -p 3000:3000 -it nobones-api

   ____    __
  / __/___/ /  ___
 / _// __/ _ \/ _ \
/___/\__/_//_/\___/ v4.6.1
High performance, minimalist Go web framework
https://echo.labstack.com
____________________________________O/_______
                                    O\
⇨ http server started on [::]:3000

```