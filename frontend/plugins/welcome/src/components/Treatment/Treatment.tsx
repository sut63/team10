import React, { useEffect, FC } from 'react';
import { Content, Header, Page, pageTheme, Link, } from '@backstage/core';
import { makeStyles, Theme, createStyles } from '@material-ui/core/styles';
import { Grid, MenuItem, Button, TextField, FormControl, Select, Typography } from '@material-ui/core';

import { Alert } from '@material-ui/lab';
import Paper from '@material-ui/core/Paper';
import { Link as RouterLink } from 'react-router-dom';

import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import { Avatar } from '@material-ui/core';
import { Cookies } from 'react-cookie/cjs';//cookie


import { DefaultApi } from '../../api/apis';
import { EntUser } from '../../api/models/EntUser';
import { EntTypetreatment } from '../../api/models/EntTypetreatment';
import { EntDoctor } from '../../api/models/EntDoctor';
import { EntPatientrecord } from '../../api/models/EntPatientrecord';
import { EntTreatment } from '../../api/models/EntTreatment';

// header css
const HeaderCustom = {
    minHeight: '50px',
};

const cookies = new Cookies();
const Name = cookies.get('Name');
const Img = cookies.get('Img');

const useStyles = makeStyles((theme: Theme) =>
    createStyles({
        button: {
            display: 'block',
            marginTop: theme.spacing(2),
        },
        formControl: {
            width: 400,
        },
        selectEmpty: {
            marginTop: theme.spacing(1),
        },
        paper: {
            padding: theme.spacing(2),
            textAlign: 'center',
            color: theme.palette.text.secondary,
        },
        table: {
            minWidth: 650,
        },
    }),
);
export default function ComponentsTable() {
    const classes = useStyles();
    const http = new DefaultApi();

    const [status, setStatus] = React.useState(false);
    const [alert, setAlert] = React.useState(true);
    const [loading, setLoading] = React.useState(true);

    const [typetreatments, setTypetreatments] = React.useState<EntTypetreatment[]>([]);
    const [doctors, setDoctors] = React.useState<EntDoctor[]>([]);
    const [patientrecords, setPatientrecords] = React.useState<EntPatientrecord[]>([]);
    const [treatments, setTreatments] = React.useState<EntTreatment[]>([]);
    const [treatmentid, setTreatmentId] = React.useState(Number);
    const [treatment, setTreatment] = React.useState(String);
    const [Users, setUsers] = React.useState<Partial<EntUser>>();

    useEffect(() => {
        const getDocdor = async () => {
            const res = await http.listDoctor({ limit: 100, offset: 0 });
            setLoading(false);
            setDoctors(res);
        };
        const getTypetreatment = async () => {
            const res = await http.listTypetreatment({ limit: 3, offset: 0 });
            setLoading(false);
            setTypetreatments(res);
        };
        const getPatientrecord = async () => {
            const res = await http.listPatientrecord({ limit: 2, offset: 0 });
            setLoading(false);
            setPatientrecords(res);
        };
        const getTreatment = async () => {
            const res = await http.listTreatment({ limit: 100, offset: 0 });
            setLoading(false);
            setTreatments(res);
        };
        const getImg = async () => {
            const res = await http.getUser({ id: Number(Img) });
            setLoading(false);
            setUsers(res);
        };
        getPatientrecord();
        getTypetreatment();
        getTreatment();
        getDocdor();
        getImg();
    }, [loading]);
    return (
        <div>

            <Page theme={pageTheme.home}>
                <Header style={HeaderCustom} title={`Treatment Department`}>
                    <Avatar alt="Remy Sharp" src={Users?.images as string} />
                    <div style={{ marginLeft: 10 }}>{Name}</div>
                </Header>
                <Content>
                    <Grid container spacing={3} >
                        <Grid container item xs={12} sm={12}  >
                            <Grid item xs={3}>
                                <Paper >
                                    <Typography align="center">
                                        <Typography align="center" variant="subtitle1">
                                            <br />เลขที่การรักษา : {treatmentid}
                                            <br />รายละเอียดการรักษา<br />
                                            <TableCell align="center">{treatment}</TableCell>
                                        </Typography>
                                    </Typography>
                                    <br />
                                    <Typography align="center">
                                        <Link component={RouterLink} to="/">
                                            <Button variant="contained" color="primary" style={{ backgroundColor: "#21b6ae" }}>
                                                กลับสู่หน้าหลัก
            </Button>
                                        </Link>
                                    </Typography>
                                    <br />
                                </Paper>
                            </Grid>

                            <Grid item xs={9}>
                                <Paper>
                                    <TableContainer component={Paper}>

                                        <Table className={classes.table} aria-label="simple table">
                                            <TableHead>
                                                <TableRow>
                                                    <TableCell align="center">เลขที่การรักษา</TableCell>
                                                    <TableCell align="center">แพทย์</TableCell>
                                                    <TableCell align="center">ผู้เข้ารับการรักษา</TableCell>
                                                    <TableCell align="center">รูปแบบการรักษา</TableCell>
                                                    <TableCell align="center">วันเวลาที่รักษา</TableCell>                                                <TableCell align="center">
                                                        <Link component={RouterLink} to="/createTreatment">
                                                            <Button variant="contained" color="primary" style={{ backgroundColor: "#21b6ae" }}>
                                                                บันทึกการรักษา
            </Button>
                                                        </Link></TableCell>
                                                </TableRow>
                                            </TableHead>
                                            <TableBody>
                                                {treatments.map(item => (doctors.filter(t => t.id === item.edges?.edgesOfDoctor?.id).map(item2 => (
                                                    <TableRow key={item.id}>
                                                        <TableCell align="center">{item.id}</TableCell>
                                                        <TableCell align="center">{item2.edges?.edgesOfDoctorinfo?.doctorname} {item2.edges?.edgesOfDoctorinfo?.doctorsurname}</TableCell>
                                                        <TableCell align="center">{item.edges?.edgesOfPatientrecord?.name}</TableCell>
                                                        <TableCell align="center">{item.edges?.edgesOfTypetreatment?.typetreatment}</TableCell>
                                                        <TableCell align="center">{item.datetreat}</TableCell>
                                                        <TableCell align="center">
                                                            <Button
                                                                onClick={() => {
                                                                    setTreatment(item.treatment as string);
                                                                    setTreatmentId(item.id as number);
                                                                }}
                                                                variant="outlined"
                                                                color="primary"
                                                            >
                                                                ดูรายละเอียด
                              </Button>
                                                        </TableCell>
                                                    </TableRow>
                                                ))))}
                                            </TableBody>
                                        </Table>
                                    </TableContainer>
                                </Paper>
                            </Grid>
                        </Grid>
                    </Grid>
                </Content>
            </Page>
        </div>
    );
};
