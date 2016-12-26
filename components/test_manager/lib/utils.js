'user strict';

const _ = require('ramda')

// traverse :: (Applicative f, Traversable t) => (a -> f a) -> (b -> f c) -> t a -> f ({succeed: t c, failed: t c})
const traverse = _.curry((of, f, traversable) => {
  return traversable
    .map(f)
    .reduce((a, b) => {
      return a.chain(y => {
        const success = x => {y.succeed.push(x); return y;};
        const failure = x => {y.failed.push(x); return y;};
        return b.fold(failure, success);
      });
    }, of({succeed: [], failed: []}));
});

// convert :: {a} -> [{ k :: String, v :: a }]
const convert = _.compose(_.map(_.zipObj(['k', 'v'])), _.toPairs);

const trace = _.curry(x => {
	console.log('Inspect::');
	console.log(x);
	console.log('')
	return x;
});

module.exports.convert = convert;
module.exports.traverse = traverse;
module.exports.trace = trace;