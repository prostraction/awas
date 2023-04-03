import React, {Component, Fragment} from "react";
import { BrowserRouter as Rounter, Routes, Route, Link, useParams } from "react-router-dom";

import Home from './components/Home'
import Admin from './components/Admin'
import TrackNumber from './components/TrackNumber'

export default class App extends Component {
    render() {
        return (
            <Rounter>
            <div className="container">
                <div className="row">
                    <h1 className="mt-1">
                        Test
                    </h1>
                    <hr className="mb-3"></hr>
                </div>

                <div className="row">
                    <div className="col-md-2">
                        <nav>
                            <ul className="list-group">
                                <li className="list-group-item">
                                    <Link to="/">Home</Link>
                                </li>
                                <li className="list-group-item">
                                    <Link to="/admin">Admin</Link>
                                </li>
                                <li className="list-group-item">
                                    <Link to="/track-number">Track number</Link>
                                </li>
                            </ul>
                        </nav>
                    </div>

                    <div className="col-md-10">
                        <Routes>
                            <Route path='/' element={<Home/>}/>
                            <Route path='/admin' element={<Admin/>}/>
                            <Route path='/track-number/:id' element={<Number/>}/>
                            <Route path='/track-number' element={<TrackNumber/>}/>
                        </Routes>
                    </div>
                </div>
            </div>
            </Rounter>
        )
    }
}

function Number() {
    let {id} = useParams();
    return <h2>Number id = {id} </h2>
}