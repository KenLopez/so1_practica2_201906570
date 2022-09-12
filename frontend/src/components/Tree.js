import React, { useEffect, useState } from 'react';
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
  const [data, setData] = useState(props.rows.map(item => {
    item.parentId = item.padre
    return item;
  }));
  const [tableColumnExtensions] = useState([
    { columnName: 'nombre', width: 300 },
  ]);

  useEffect(() => {
    const update = () => {
        setData(props.rows.map(item => {
            item.parentId = item.padre
            return item;
          })
        )
        console.log(data);
    }
    update()
  }, [props.rows])
  

  return (
    <div>
      <Grid
        rows={data}
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
          for="name"
        />
      </Grid>
    </div>
  );
};


export default Tree