package clb

type DescribeForwardLBListenersArgs struct {
	LoadBalancerId   string    `qcloud_arg:"loadBalancerId,required"`
	ListenerIds      *[]string `qcloud_arg:"listenerIds"`
	Protocol         *int      `qcloud_arg:"protocol"`
	LoadBalancerPort *int32    `qcloud_arg:"loadBalancerPort"`
}

type DescribeForwardLBtenersResponse struct {
	Response
	TotalCount  int        `json:"totalCount"`
	ListenerSet []ListenerForward `json:"listenerSet"`
}

type ListenerForward struct {
	ListenerId       string `json:"listenerId"`
	LoadBalancerPort int32  `json:"loadBalancerPort"`
	Protocol         int    `json:"protocol"`
	ProtocolType     string `json:"protocolType"`
	SSLMode          string `json:"SSLMode"`
	CertId           string `json:"certId"`
	CertCaId         string `json:"certCaId"`
	Rules            []ListenerForwardRule `json:"rules"`
}

type ListenerForwardRule struct {

}

func (client *Client) DescribeForwardLBListeners(args *DescribeForwardLBListenersArgs) (
	*DescribeForwardLBtenersResponse,
	error,
) {
	response := &DescribeForwardLBtenersResponse{}
	err := client.Invoke("DescribeForwardLBListeners", args, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

type CreateForwardLBFourthLayerListenersArgs struct {
	LoadBalancerId string               `qcloud_arg:"loadBalancerId,required"`
	Listeners      []CreateForwardLBFourthLayerListenersOpts `qcloud_arg:"listeners"`
}

type CreateForwardLBFourthLayerListenersOpts struct {
	LoadBalancerPort int32   `qcloud_arg:"loadBalancerPort,required"`
	Protocol         int     `qcloud_arg:"protocol,required"`
	ListenerName     *string `qcloud_arg:"listenerName"`
	SessionExpire    *int    `qcloud_arg:"sessionExpire"`
	HealthSwitch     *int    `qcloud_arg:"healthSwitch"`
	TimeOut          *int    `qcloud_arg:"timeOut"`
	IntervalTime     *int    `qcloud_arg:"intervalTime"`
	HealthNum        *int    `qcloud_arg:"healthNum"`
	UnhealthNum      *int    `qcloud_arg:"unhealthNum"`
	HttpHash         *int    `qcloud_arg:"httpHash"`
	Scheduler         *int    `qcloud_arg:"scheduler"`
}

type CreateForwardLBFourthLayerListenersResponse struct {
	Response
	RequestId   int      `json:"requestId"`
	ListenerIds []string `json:"listenerIds"`
}

func (response CreateForwardLBFourthLayerListenersResponse) Id() int {
	return response.RequestId
}

func (client *Client) CreateForwardLBFourthLayerListeners(args *CreateForwardLBFourthLayerListenersArgs) (
	*CreateForwardLBFourthLayerListenersResponse,
	error,
) {
	response := &CreateForwardLBFourthLayerListenersResponse{}
	err := client.Invoke("CreateForwardLBFourthLayerListeners", args, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

type DeleteForwardLBListenerArgs struct {
	LoadBalancerId string   `qcloud_arg:"loadBalancerId,required"`
	ListenerId     string `qcloud_arg:"listenerId,required"`
}

type DeleteForwardLBListenerResponse struct {
	Response
	RequestId int `json:"requestId"`
}

func (response DeleteForwardLBListenerResponse) Id() int {
	return response.RequestId
}

func (client *Client) DeleteForwardLBListener(LoadBalancerId string, ListenerId string) (
	*DeleteForwardLBListenerResponse,
	error,
) {
	response := &DeleteForwardLBListenerResponse{}
	err := client.Invoke("DeleteForwardLBListener", &DeleteForwardLBListenerArgs{
		LoadBalancerId: LoadBalancerId,
		ListenerId: ListenerId,
	}, response)

	if err != nil {
		return nil, err
	}

	return response, nil
}