import React, { useState, useEffect } from 'react';
import Backdrop from '@material-ui/core/Backdrop';
import Button from '@material-ui/core/Button';
import { DefaultApi } from '../../api/apis';
import { EntHistorytaking } from '../../api/models';
import { makeStyles, Theme, createStyles } from '@material-ui/core/styles';
import { Grid, Modal, Typography } from '@material-ui/core';
import Fade from '@material-ui/core/Fade';

const useStyles = makeStyles((theme: Theme) =>
    createStyles({
        modal: {
            display: 'flex',
            alignItems: 'center',
            justifyContent: 'center',
        },
        paper: {
            backgroundColor: theme.palette.background.paper,
            border: '2px solid #000',
            boxShadow: theme.shadows[5],
            padding: theme.spacing(2, 4, 3),
        },
    }),
);

export default function MoreInfo(id: any) {
    const classes = useStyles();
    const [open, setOpen] = React.useState(false);

    const handleOpen = () => {
        setOpen(true);
    };

    const handleClose = () => {
        setOpen(false);
    };
    const http = new DefaultApi();
    const [loading, setLoading] = React.useState(true);
    const [Historytaking, setHistorytakings] = useState<EntHistorytaking[]>(Array);

    const getHistorytakings = async () => {
        const res = await http.listHistorytaking({ offset: 0 })
        setLoading(false);
        setHistorytakings(res);
        console.log(id.id)
    };
    useEffect(() => {
        getHistorytakings();

    }, [loading]);

    return (
        <div>
            <Button variant='outlined' color='primary' onClick={handleOpen}>
                รายละเอียดเพิ่มเติม
            </Button>
            <Modal
                aria-labelledby="transition-modal-title"
                aria-describedby="transition-modal-description"
                className={classes.modal}
                open={open}
                onClose={handleClose}
                closeAfterTransition
                BackdropComponent={Backdrop}
            >
                <Fade in={open}>
                    <div className={classes.paper}>
                    {Historytaking.filter(h => h.id === id.id).map(historytaking =>( 
                        <Grid>
                            <Typography variant='h6'>
                                รายละเอียด<br />
                            </Typography>
                            <Typography variant='subtitle1'>
                                Patientname:&emsp;{historytaking.edges?.edgesOfPatientrecord?.name}<br/>
                                Nurse:&emsp;{historytaking.edges?.edgesOfNurse?.name}<br/>
                                Department:&emsp;{historytaking.edges?.edgesOfDepartment?.department}<br/>
                                Symptomseverity:&emsp;{historytaking.edges?.edgesOfSymptomseverity?.symptomseverity}<br/>
                                Hight:&emsp;{historytaking.hight}&emsp;cm<br />
                                Weight:&emsp;{historytaking.weight}&emsp;kg<br />
                                Temperature:&emsp;{historytaking.temp}&emsp;°C <br />
                                Pulse:&emsp;{historytaking.pulse}&emsp;times/min<br />
                                Respiration:&emsp;{historytaking.respiration}&emsp;times/min<br />
                                Blood pressure:&emsp;{historytaking.bp}&emsp;mm/Hg<br />
                                %Oxygen:&emsp;{historytaking.oxygen}<br />
                                Sympton:&emsp;{historytaking.symptom}<br />
                                Datetime:&emsp;{historytaking.datetime}<br />
                            </Typography>
                        </Grid>
                    ))}
                    </div>
                </Fade>
            </Modal>
        </div>
    );
}