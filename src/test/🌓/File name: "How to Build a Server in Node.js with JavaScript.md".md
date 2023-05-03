Here is the JavaScript code for building a basic server using Node.js and the Express package:

```
// Require the Express package
const express = require('express');

// Create a new Express app
const app = express();

// Define a route for the homepage
app.get('/', (req, res) => {
  res.send('Hello World!');
});

// Start the server on port 3000
app.listen(3000, () => {
  console.log('Server listening on port 3000');
});
```

To run this code, you would save it in a file called `server.js` and execute the following command in your terminal:

```
node server.js
```

This would start the server and allow you to access it by visiting `http://localhost:3000` in a web browser.