package goeasypost
import "fmt"

func NewShipment(reqShipment *Shipment) (*Shipment, error) {

	params := make(map[string]string)

	err := flattenStructMap(reqShipment, "shipment", params)
	if err != nil {
		return nil, err
	}

	var shipment = new(Shipment)

	err = createEntity("shipments", params, shipment)
	if err != nil {
		return nil, err
	}

	return shipment, nil
}

func GetShipment(id string) (*Shipment, error) {

	shipment := &Shipment{}

	err := getEntityById("shipments", id, shipment)
	if err != nil {
		return nil, err
	}

	return shipment, nil
}

func BuyShipment(shipmentId string, rateId string, insurance float32) (*Shipment, error) {

	var params = make(map[string]string)

	params["rate[id]"] = rateId
	params["insurance"] = fmt.Sprintf("%s",insurance)


	shipment := &Shipment{}

	err := callEntityWithCommand("shipments", shipmentId,"buy",params, shipment)
	if err != nil {
		return nil, err
	}

	return shipment, nil
}

func InsureShipment(shipmentId string, insurance float32) (*Shipment, error) {

	var params = make(map[string]string)

	params["amount"] = fmt.Sprintf("%s",insurance)


	shipment := &Shipment{}

	err := callEntityWithCommand("shipments", shipmentId,"insure",params, shipment)
	if err != nil {
		return nil, err
	}

	return shipment, nil
}

func RefundShipment(shipmentId string) (*Shipment, error) {


	shipment := &Shipment{}

	err := callEntityWithCommand("shipments", shipmentId,"refund",nil, shipment)
	if err != nil {
		return nil, err
	}

	return shipment, nil
}
