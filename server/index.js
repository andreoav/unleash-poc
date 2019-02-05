const unleash = require('unleash-server');

unleash.start({
	databaseUrl: 'postgres://unleash:unleash@localhost:5432/unleash',
	enableLegacyRoutes: false,
	preRouterHook: app => {
		app.use('/api/client', (req, res, next) => {
			if (req.header('authorization') !== '12345') {
				res.sendStatus(401);
			} else {
				next();
			}
		});
	},
});
