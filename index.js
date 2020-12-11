const express = require('express');
const cors = require('cors');
const app = express();
const port = process.env.PORT || 3000;

app.use(cors());

app.get('/', (req, res) => {
    res.json({
        utc_datetime: new Date().toISOString(),
    });
});

const listener = app.listen(port, () => {
    console.log(`Time service app listening at ${listener.address().port}`);
});

