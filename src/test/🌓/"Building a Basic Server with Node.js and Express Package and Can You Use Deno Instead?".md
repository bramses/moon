Yes, you can use Deno instead of Node.js to create a web server using its built-in web server API. Here is the code for building a server in Deno:

```
// Import the HTTP module
import { serve } from "https://deno.land/std/http/server.ts";

// Create a new server
const server = serve({ port: 3000 });

// Declare a listener for incoming requests
for await (const req of server) {
  // Return a response
  req.respond({ body: "Hello World!" });
}
```

To run the server, you would save the code in a file called `server.ts` and execute the following command in your terminal:
```
deno run --allow-net server.ts
```

This would start the server and allow you to access it by visiting `http://localhost:3000` in a web browser.