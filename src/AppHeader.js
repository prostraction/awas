import React, { Component, Fragment } from "react";
//import './AppHeader.css'

export default class AppHeader extends Component {

    constructor(props) {
        super(props);
        this.handlePostChange = this.handlePostChange.bind(this);
    }

    handlePostChange(posts) {
        this.props.handlePostChange(posts);
    }

    render() {
        return (
            <Fragment>
            <h1>{this.props.title}</h1>
            <p>=== {this.props.posts.length} ===</p>
            </Fragment>
        );
    }
}