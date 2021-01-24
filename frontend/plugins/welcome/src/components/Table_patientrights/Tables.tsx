import React, { useState, useEffect } from 'react';
import { makeStyles } from '@material-ui/core/styles';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';

import Button from '@material-ui/core/Button';
import { DefaultApi } from '../../api/apis';
import { EntPatientrights } from '../../api/models/EntPatientrights';

const useStyles = makeStyles(theme => ({
  table: {
    minWidth: 650,
  },
  formControl: {
    margin: theme.spacing(3),
    width: 200,
  },
}));



export default function ComponentsTable(sim: any) {
  const classes = useStyles();
  const api = new DefaultApi();
  const [Patientrightss, setPatientrightss] = useState<EntPatientrights[]>(Array);




  const [loading, setLoading] = useState(true);

  const getPatientrightss = async () => {
    const res = await api.listPatientrights({ limit: 100, offset: 0 });
    setLoading(false);
    setPatientrightss(res);
  };


  //console.log(Patientrightss)


  useEffect(() => {

    getPatientrightss();


  }, [loading]);










  const deletePatientrightsss = async (id: number) => {
    const res = await api.deletePatientrights({ id: id });
    setLoading(true);
  };
  var p = 0;
  console.log(sim.sim)
  for (var val of Patientrightss) {
    if (val.edges?.edgesOfPatientrightsPatientrecord?.id === sim.sim || sim.sim === 0) {
      p = p + 1
    }
    //console.log(val)
  }
  console.log("p = ", p)
  if (p !== 0) {
    return (

      <TableContainer component={Paper}>
        <Table className={classes.table} aria-label="simple table">
          <TableHead>
            <TableRow>
              <TableCell align="center">No.</TableCell>
              <TableCell align="center">ชื่อ ประกัน</TableCell>
              <TableCell align="center">Medicalrecordstaff</TableCell>
              <TableCell align="center">ความสามารถสิทธิ์</TableCell>
              <TableCell align="center">ผู้ป่วย</TableCell>
              <TableCell align="center">ผู้ดูแล</TableCell>
              <TableCell align="center">วันที่สร้าง(ปี-เดือน-วัน)</TableCell>
              <TableCell align="center">Manage</TableCell>

            </TableRow>
          </TableHead>
          <TableBody>
            {Patientrightss === undefined
              ? null
              : Patientrightss.filter(i => i.edges?.edgesOfPatientrightsPatientrecord?.id === sim.sim || sim.sim === 0).map((item: any) => (
                <TableRow>
                  <TableCell align="center">{item.id}</TableCell>
                  <TableCell align="center">{item.edges?.edgesOfPatientrightsInsurance?.insurancecompany}</TableCell>
                  <TableCell align="center">{item.edges?.edgesOfPatientrightsMedicalrecordstaff?.name}</TableCell>

                  <TableCell align="center">{item.edges?.edgesOfPatientrightsAbilitypatientrights?.check}</TableCell>
                  <TableCell align="center">{item.edges?.edgesOfPatientrightsPatientrecord?.name}</TableCell>
                  <TableCell align="center">{item.responsible}</TableCell>
                  <TableCell align="center">{item.permissionDate}</TableCell>
                  <TableCell align="center">
                    <Button
                      onClick={() => {
                        deletePatientrightsss(item.id);
                      }}
                      style={{ marginLeft: 10 }}
                      variant="contained"
                      color="secondary"
                    >
                      Delete
               </Button>
                  </TableCell>




                </TableRow>
              ))}
          </TableBody>
        </Table>
      </TableContainer>
    );
  } else {
    return (
      <TableContainer component={Paper}>
        <Table className={classes.table} aria-label="simple table">
          <TableHead>
            <TableRow>
              <TableCell align="center">ไม่พบสิทธ์</TableCell>
            </TableRow>

          </TableHead>

        </Table>
      </TableContainer>
    );
  }



}

