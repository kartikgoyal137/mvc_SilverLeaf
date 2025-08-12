
import Navbar from "../components/Navbar"
import HeroImg from '../assets/cafe.jpg'
import Chef from '../assets/chef.png'
import Table from '../assets/table.png'
import './css/home.css'
import { useNavigate } from "react-router-dom"
import Review from '../components/ReviewCard'

export default function Home() {
    const nav = useNavigate()
    function Menu() {
        nav('/menu')
    }

    return (
        <>
        <Navbar />
        <div className="container-fluid hero-bg" style={{ backgroundImage: `url(${HeroImg})`, height: "100vh" }}>
                
                <div className="hero-overlay h-100"></div>
                <div className="container text-center d-flex flex-column align-items-center justify-content-center h-75">
                    <h1 className="pt-5 cormorant heading display-1 fw-bold w-100 w-lg-75">Experience Culinary Excellence</h1>
                    <p className="fs-3 pt-4 w-100 w-lg-75">A sanctuary for the senses, where every dish is a masterpiece and every moment is cherished to create a unique experience.</p>
                    <button onClick={Menu} className="btn btn-warning fs-5 fw-bold mt-4">Book a table</button>
                </div>
            </div>

            <div className="container py-5">
                <div className="row align-items-center">
                    <div className="col-lg-7">
                        <div className="story">
                            <p className="fs-2 cormorant mb-0">Our Story</p>
                            <h3 className="cormorant display-3 mb-4 mt-0 fw-semibold">The Heart of Silver Leaf</h3>
                            <p className="fs-4">Founded in 2010, Silver Leaf began as a humble dream to bring authentic, high-quality cuisine to the forefront of our city's dining scene. Our philosophy is simple: use the freshest, locally-sourced ingredients to create unforgettable dishes that tell a story. Over the years, we've grown into a beloved establishment, but our core values of quality, community, and culinary passion remain unchanged.</p>
                        </div>
                    </div>
                    <div className="col-lg-5 mt-4 mt-lg-0">
                        <img src={Chef} alt="Chef in the kitchen" className="img-fluid rounded-5 mb-3" />
                        <img src={Table} alt="Restaurant dining table" className="img-fluid rounded-5" />
                    </div>
                </div>
            </div>

            <div className="container-fluid section-3 pt-5 pb-5">
                <div className="container d-flex flex-column align-items-center">
                    <p className="cormorant text-center fs-2 mb-0">Taste Our Passion</p>
                    <h1 className="cormorant text-center display-3">Featured Dishes</h1>
                    <p className="text-center fs-5 mt-3 w-100 w-lg-50 mb-5">From our kitchen to your table, these are the creations that define us. Handpicked by our chef, celebrated by our guests.</p>
                </div>
                <div className="container">
                    <div className="row justify-content-center text-center">
                        <div className="col-12 col-md-6 col-lg-4 mb-4">
                            <div className="food">
                                <img alt="Seared Scallops" className="img-fluid rounded-5 mb-4" src="https://lh3.googleusercontent.com/aida-public/AB6AXuDK_fnzBrvq62OosPBTLExYT2_pQRZEMsiaaqKQHusFd_GqLsXIPHF6mVEGdr5AiWUQSvBdL1d_oBK1rb3Xh-zVkgnDjsmI7Zy7F6K1pPJ81vKFkZg04hv35CjgE_-6cZvhyOYHQ8X1pfXxvBu15OAZ14XwFr7p-VeXHblq2BocWLBJi_-UwXQq8Q24usqNLBT9-jkfiOS2o7s1g_iyqTq0Ow3_hMHNy-cHXOpXzXw6wKz7e82L2pCN2AgO1wAwDj-uOcW0aeFoQ9g" />
                                <p className="fw-bold fs-3 cormorant">Pan-Seared Scallops</p>
                                <p className="fs-6">With saffron risotto and asparagus tips.</p>
                            </div>
                        </div>
                        <div className="col-12 col-md-6 col-lg-4 mb-4">
                            <div className="food">
                                <img alt="Wagyu Steak" className="img-fluid rounded-5 mb-4" src="https://lh3.googleusercontent.com/aida-public/AB6AXuA6bQrTJUHU3--2vy9brvpoNIAgFjFUqsPGlw_VdyfeQHyj177TBgPJ2a5DzugAAIS4Wd-Iz18Rdgx18JkHeXV4_p7XASJEDXSraBGOXHtSSWItDIFGUSx3-cAS7UI8Ij_aFZAGz_mu9ZAezwVMWbdO8meE5t6QI43wS_t73noN-KGREJ4AwSKheG1kZ4mAviLc-9dZQRD9yZ9iTugAzOBmi1N1Cqw3UrIZHsBSgtJ3wktuM2J9Gt6SncaKDTobGfnLEl1s6evV67g" />
                                <p className="fw-bold fs-3 cormorant">Prime Wagyu Steak</p>
                                <p className="fs-6">Served with truffle mashed potatoes and a red wine reduction.</p>
                            </div>
                        </div>
                        <div className="col-12 col-md-6 col-lg-4 mb-4">
                            <div className="food">
                                <img alt="Truffle Pasta" className="img-fluid rounded-5 mb-4" src="https://lh3.googleusercontent.com/aida-public/AB6AXuDKkKlAAk4IEaWqeRs1Ld4MUNUmYlP2xoOf-0NsUnJuPybDXGtmu-Vmwkn4lBCcKWk4xTkZWJlKsfRiQKqYPp2_vvqx4q6SA-1y1-Jktm9-2nL-F1KU8RxlAqn9avoyLM4oSH5hj1QoDBYJbNaIDvtf7JmBKbFW11wCuphelsnNUazwi70pFC-H8m_QWzlVNAOeTNIRLp9qUrb7wHJjHkoDxfeLztU6mkLRWDy1fRfPhvBU6baTLtrCpoNVljFXHcvhmGWkwlM3aWE" />
                                <p className="fw-bold fs-3 cormorant">Black Truffle Pasta</p>
                                <p className="fs-6">Handmade tagliatelle in a creamy parmesan and black truffle sauce.</p>
                            </div>
                        </div>
                        
                    </div>
                </div>
            </div>

            {/* REVIEWS SECTION - Restructured with Bootstrap Grid */}
            <div className="container my-5">
                <div className="row text-center">
                    {/* Each review gets a column that stacks on mobile */}
                    <div className="col-12 col-md-4 mb-4 mb-md-0">
                        <Review />
                    </div>
                    <div className="col-12 col-md-4 mb-4 mb-md-0">
                        <Review />
                    </div>
                    <div className="col-12 col-md-4">
                        <Review />
                    </div>
                </div>
            </div>
        </>
    )
}