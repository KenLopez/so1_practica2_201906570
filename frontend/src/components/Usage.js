import React, { useState } from "react";
import { Container } from "react-bootstrap";
import { Chart as ChartJS, ArcElement, Tooltip, Legend } from 'chart.js';
import { Doughnut } from 'react-chartjs-2';

ChartJS.register(ArcElement, Tooltip, Legend);

function Usage(props) {
    const data = {
        labels: ['En uso', 'Libre'],
        datasets: [
            {
            label: 'Porcentaje',
            data: [props.usage, 100-props.usage],
            backgroundColor: [
                'rgba(255, 99, 132, 0.2)',
                'rgba(75, 192, 192, 0.2)',
            ],
            borderColor: [
                'rgba(255, 99, 132, 1)',
                'rgba(75, 192, 192, 1)',
            ],
            borderWidth: 1,
            },
        ],
    };
  return (
    <Container fluid>
        <h3>{props.title}</h3>
        <hr/>
        <Doughnut data={data} />
    </Container>
  );
}

export default Usage;