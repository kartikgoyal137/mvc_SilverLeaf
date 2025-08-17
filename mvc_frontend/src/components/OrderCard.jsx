const FormattedDate = ({ isoDateString }) => {
  if (!isoDateString) return null;
  const options = {
    year: 'numeric',
    month: 'long',
    day: 'numeric',
    hour: '2-digit',
    minute: '2-digit'
  };
  const date = new Date(isoDateString);
  return new Intl.DateTimeFormat('en-US', options).format(date);
};

const OrderCard = ({ order }) => {
    
    const getStatusBadgeClass = (status) => {
        switch (status) {
            case 'Completed':
                return 'bg-success';
            case 'Cooking':
                return 'bg-info text-dark';
            case 'Yet to start':
                return 'bg-secondary';
            default:
                return 'bg-light text-dark';
        }
    };

    return (
      
        <div className="col-md-6 col-lg-4 mb-4">
            <div className="card h-100 shadow-sm rounded-4">
                <div className="card-header d-flex justify-content-between align-items-center bg-light-subtle rounded-top-4">
                    <h5 className="mb-0 fw-bold">Order #{order.order_id}</h5>
                    <span className={`badge ${getStatusBadgeClass(order.status)} p-2`}>{order.status}</span>
                </div>
                <div className="card-body d-flex flex-column">
                    <h6 className="card-subtitle mb-3 text-muted">
                        <FormattedDate isoDateString={order.created_at} />
                    </h6>
                   
                    <ul className="list-group list-group-flush mb-3">
                        {order.products.map((product, index) => (
                            <li key={index} className="list-group-item d-flex justify-content-between align-items-center px-0">
                                {product.name}
                                <span className="badge bg-primary rounded-pill">Qty: {product.quantity}</span>
                            </li>
                        ))}
                    </ul>
                 
                    {order.instructions && (
                        <div className="mt-auto">
                            <strong>Instructions:</strong>
                            <p className="card-text fst-italic bg-light p-2 rounded">"{order.instructions}"</p>
                        </div>
                    )}
                </div>
                <div className="card-footer text-muted bg-light-subtle rounded-bottom-4">
                    Table No: <strong>{order.table_no || 'N/A'}</strong>
                </div>
            </div>
        </div>
    );
};

export default OrderCard;