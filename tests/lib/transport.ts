const transport = endpoint => async <T>(method: string, args: any = []) => {
	return new Promise<T>(async (resolve, reject) => {
		fetch(
			`${endpoint}/${encodeURIComponent(method)}`,
			{
				method: 'POST',
				body: JSON.stringify(args),
			}
		)
		.then((response) => {
			return response.json() as Promise<T>;
		})
		.then((val) => resolve(val))
		.catch((err) => reject(err))
	})
};

export default transport
