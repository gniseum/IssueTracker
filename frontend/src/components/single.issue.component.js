import React, {useState, useEffect} from 'react';
import 'bootstrap/dist/css/bootstrap.css';
import { Button, Card, Row, Col } from 'react-bootstrap';

const Issue = ({issueData, setChangeTaskedUser, deleteSingleIssue, setChangeIssue}) => {
    return (

    )

    function changeTaskedUser() {
        setChangeTaskedUser(
            {
            "change": true,
            "id": issueData._id
            }
        )
    }

    function changeIssue() {
        setChangeIssue(
            {
            "change": true,
            "id": issueData._id
            }
        )
    }
}

export default Issue
