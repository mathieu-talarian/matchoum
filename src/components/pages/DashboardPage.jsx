import React from "react";
import {connect} from "react-redux";
import PropTypes from "prop-types";
import {Redirect} from "react-router-dom";

import ConfirmEmailMessage from "../messages/ConfirmEmailMessage";

const DashboardPage = ({isConfirmed, hasInfos}) => (
    <div>
        {!isConfirmed && <ConfirmEmailMessage/>}
        {!hasInfos && isConfirmed && <Redirect to="/profil"/>}
    </div>
);

DashboardPage.propTypes = {
    isConfirmed: PropTypes.bool.isRequired,
    hasInfos: PropTypes.bool.isRequired
};

function mapStateToProps(state) {
    return {
        isConfirmed: !!state.user.confirmed,
        hasInfos: state.user.infos
    };
}

export default connect(mapStateToProps)(DashboardPage);
