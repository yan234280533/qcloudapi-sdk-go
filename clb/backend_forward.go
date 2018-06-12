package clb

type DescribeForwardLBBackendsArgs struct {
	LoadBalancerId string `qcloud_arg:"loadBalancerId,required"`
	Listeners      []string `qcloud_arg:"listeners"`
	Protocol        int    `qcloud_arg:"protocol"`
	LoadBalancerPort  int    `qcloud_arg:"loadBalancerPort"`
}

type DescribeForwardLBBackendsResponse struct {
	Response
	Data []ForwardLBBackendsData `json:"data"`
}

type ForwardLBBackendsData struct {
	ListenerID       string `json:"listenerId"`
	LoadBalancerPort int    `json:"loadBalancerPort"`
	Protocol         int    `json:"protocol"`
	ProtocolType     string `json:"protocolType"`
	Backends   []ForwardLBBackendsInfo`json:"backends"`
}

type ForwardLBBackendsInfo struct {
	AddTimestamp   string   `json:"addTimestamp"`
	InstanceName   string   `json:"instanceName"`
	InstanceStatus int      `json:"instanceStatus"`
	LanIP          string   `json:"lanIp"`
	Port           int      `json:"port"`
	UnInstanceID   string   `json:"unInstanceId"`
	WanIPSet       []string `json:"wanIpSet"`
	Weight         int      `json:"weight"`
}

func (client *Client) DescribeForwardLBBackends(loadBalancerId string, listenerIds []string,protocol int, loadBalancerPort int) (
	*DescribeForwardLBBackendsResponse,
	error,
) {
	args := &DescribeForwardLBBackendsArgs{
		LoadBalancerId: loadBalancerId,
		Listeners: listenerIds,
		Protocol: protocol,
		LoadBalancerPort: loadBalancerPort,
	}

	response := &DescribeForwardLBBackendsResponse{}
	err := client.Invoke("DescribeForwardLBBackends", args, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

type RegisterInstancesWithForwardLBFourthListenerArgs struct {
	LoadBalancerId    string                  `qcloud_arg:"loadBalancerId,required"`
	ListenerId        string                  `qcloud_arg:"listenerId,required"`
	Backends       []ForwardRegisterInstancesOpts `qcloud_arg:"backends,required"`
}

type ForwardRegisterInstancesOpts struct {
	InstanceId string `qcloud_arg:"instanceId,required"`
	Port       *int   `qcloud_arg:"port"`
	Weight     *int   `qcloud_arg:"weight"`
}

type RegisterInstancesWithForwardLBFourthListenerResponse struct {
	Response
	RequestId int `json:"requestId"`
}

func (response RegisterInstancesWithForwardLBFourthListenerResponse) Id() int {
	return response.RequestId
}

func (client *Client) RegisterInstancesWithForwardLBFourthListener(args *RegisterInstancesWithForwardLBFourthListenerArgs) (
	*RegisterInstancesWithForwardLBFourthListenerResponse,
	error,
) {
	response := &RegisterInstancesWithForwardLBFourthListenerResponse{}
	err := client.Invoke("RegisterInstancesWithForwardLBFourthListener", args, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

type DeregisterInstancesFromForwardLBFourthListenerArgs struct {
	LoadBalancerId    string              `qcloud_arg:"loadBalancerId,required"`
	ListenerId        string                  `qcloud_arg:"listenerId,required"`
	Backends       []ForwardDeregisterInstances `qcloud_arg:"backends,required"`
}

type ForwardDeregisterInstances  struct {
	InstanceId string `qcloud_arg:"instanceId"`
	Port       *int   `qcloud_arg:"port"`
}

type DeregisterInstancesFromForwardLBFourthListenerResponse struct {
	Response
	RequestId int `json:"requestId"`
}

func (response DeregisterInstancesFromForwardLBFourthListenerResponse) Id() int {
	return response.RequestId
}

func (client *Client) DeregisterInstancesFromForwardLBFourthListener(LoadBalancerId string, listenId string,InstanceIds []string,port int) (
	*DeregisterInstancesFromForwardLBFourthListenerResponse,
	error,
) {

	backends := []ForwardDeregisterInstances{}
	for _, instanceId := range InstanceIds {
		backends = append(backends, ForwardDeregisterInstances{InstanceId: instanceId,Port:&port})
	}

	args := &DeregisterInstancesFromForwardLBFourthListenerArgs{
		LoadBalancerId: LoadBalancerId,
		ListenerId: listenId,
		Backends:       backends,
	}
	response := &DeregisterInstancesFromForwardLBFourthListenerResponse{}
	err := client.Invoke("DeregisterInstancesFromForwardLBFourthListener", args, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}
