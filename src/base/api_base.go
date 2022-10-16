package base

type Api struct{}

func (api Api) Get(endpoint string) NetClient {
	return HttpService().Get().Url(GetBaseUrl()+endpoint).AddHeader("Content-Type", "application/json")
}

func (api Api) Post(endpoint string) NetClient {
	return HttpService().Post().Url(GetBaseUrl()+endpoint).AddHeader("Content-Type", "application/json")
}

func (api Api) Put(endpoint string) NetClient {
	return HttpService().Put().Url(GetBaseUrl()+endpoint).AddHeader("Content-Type", "application/json")
}

func (api Api) Delete(endpoint string) NetClient {
	return HttpService().Delete().Url(GetBaseUrl()+endpoint).AddHeader("Content-Type", "application/json")
}
