export default function Navbar() {

    return (    
        <>
        <nav className="navbar navbar-expand-lg ">
            <div className="container-fluid">
                <a className="navbar-brand" href="#">Silver Leaf</a>
                <button className="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbarSupportedContent" aria-controls="navbarSupportedContent" aria-expanded="false" aria-label="Toggle navigation">
                <span className="navbar-toggler-icon"></span>
                </button>
                <div className="collapse navbar-collapse" id="navbarSupportedContent">
                <button className="btn btn-outline-success mx-2" type="submit">Login</button>
                <button className="btn btn-outline-success mx-2" type="submit">Signup</button>
            </div>
        </div>
        </nav>
        </>
    )
}