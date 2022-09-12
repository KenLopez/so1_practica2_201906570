import React, { useState } from 'react';
import {
  TreeDataState,
  CustomTreeData,
} from '@devexpress/dx-react-grid';
import {
  Grid,
  Table,
  TableHeaderRow,
  TableTreeColumn,
} from '@devexpress/dx-react-grid-bootstrap4';
import '@devexpress/dx-react-grid-bootstrap4/dist/dx-react-grid-bootstrap4.css';
import { Container } from 'react-bootstrap';

const getChildRows = (row, rootRows) => {
  const childRows = rootRows.filter(r => r.parentId === (row ? row.id : null));
  return childRows.length ? childRows : null;
};

function Tree(props) {
  const [columns] = useState([
    { name: 'pid', title: 'PID' },
    { name: 'nombre', title: 'Nombre' },
    { name: 'usuario', title: 'Usuario' },
    { name: 'ram', title: '%Ram' },
    { name: 'estado', title: 'Estado' }
  ]);
  const [tableColumnExtensions] = useState([
    { columnName: 'pid', width: 200 },
  ]);

  return (
    <Container className='card mt-5 py-3 px-5'>
        <h3>Árbol de Procesos</h3>
        <hr/>
        <Grid
            rows={props.rows}
            columns={columns}
        >
            <TreeDataState />
            <CustomTreeData
            getChildRows={getChildRows}
            />
            <Table
            columnExtensions={tableColumnExtensions}
            />
            <TableHeaderRow />
            <TableTreeColumn
            for="pid"
            />
        </Grid>
    </Container>
  );
};


export default Tree