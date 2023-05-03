Test:
1. Verify that importing the HTTP module from Deno's standard library by using the statement `import { serve } from "https://deno.land/std/http/server.ts"` is successful.
2. Verify that a new server is created by passing `{ port: 3000 }` as an argument to the `serve()` function.
3. Verify that the listener for incoming requests is declared successfully using the `for await (const req of server)` syntax.
4. Verify that the response to incoming requests is being returned successfully by using the `req.respond({ body: "Hello World!" });` statement.
5. Run the server by executing the command `deno run --allow-net server.ts` in the terminal and ensure the server starts successfully.
6. Access the server at `http://localhost:3000` in a web browser and verify that the response "Hello World!" is displayed.