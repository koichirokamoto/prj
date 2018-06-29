import * as express from 'express';
import { HELLO, GOOD_BYE } from './message';

const app = express();

app.get('/', (req: express.Request, res: express.Response) => {
  res.send(`${HELLO} from App Engine!`);
});

app.get('/bye', (req: express.Request, res: express.Response) => {
  res.send(`${GOOD_BYE} from App Engine!`);
});

const PORT = process.env.PORT || 8080;
app.listen(PORT, () => {
  console.log(`Server listening on port ${PORT}...`);
});
