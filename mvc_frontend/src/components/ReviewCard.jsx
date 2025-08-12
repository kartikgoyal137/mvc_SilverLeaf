export default function Review() {
  return (
    <div className="container py-5">
      <div className="row justify-content-center">
        <div className="col-md-4">
          <Card />
        </div>
      </div>
    </div>
  );
}

function Card() {
  return (
    <div className="card shadow-sm w-100 rounded-4">
      <img 
        src="" 
        className="card-img-top rounded-top-4" 
        alt="Placeholder"
      />
      <div className="card-body">
        <p className="card-text">
          This is the body of the card. You can put any description or content here. 
          The component is styled using Bootstrap classes to create a clean and modern look.
        </p>
      </div>
    </div>
  );
}
