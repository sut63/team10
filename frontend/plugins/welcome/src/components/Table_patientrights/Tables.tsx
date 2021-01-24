import React, { useState, useEffect } from 'react';
import { makeStyles } from '@material-ui/core/styles';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';
import TextField from '@material-ui/core/TextField';
import Button from '@material-ui/core/Button';
import { DefaultApi } from '../../api/apis';
import { EntPatientrights } from '../../api/models/EntPatientrights';

const useStyles = makeStyles({
  table: {
    minWidth: 650,
  },
});



export default function ComponentsTable() {
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



  interface se {
    pc : string
    res: string

  }
  const [Pat, setPat] = React.useState<Partial<se>>({});

  const handleChange = (
    event: React.ChangeEvent<{ name?: string; value: unknown }>,
  ) => {
    const name = event.target.name as keyof typeof ComponentsTable;
    const { value } = event.target;
    setPat({ ...Pat, [name]: value });
  };




  const deletePatientrightsss = async (id: number) => {
    const res = await api.deletePatientrights({ id: id });
    setLoading(true);
  };

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

          <TableRow>
            <TableCell align="center">
              ค้นหา
            </TableCell>
            <TableCell align="center"></TableCell>
            <TableCell align="center"></TableCell>
            <TableCell align="center"></TableCell>
            <TableCell align="center">
              <TextField
                name="pc"
                label="ผู้ป่วย"
                variant="outlined"
                type="string"
                size="medium"

                value={Pat.pc}
                onChange={handleChange}
              />
            </TableCell>
            <TableCell align="center"></TableCell>
            <TableCell align="center"></TableCell>
            <TableCell align="center">
            <Button
                    onClick={() => {
                    }}
                    style={{ marginLeft: 10 }}
                    variant="contained"
                    color="primary"
                  >
                    ค้นหา
               </Button>
            </TableCell>
          </TableRow>

        </TableBody>

        <TableBody>
          {Patientrightss === undefined
            ? null
            : Patientrightss.map((item: any) => (
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
}

