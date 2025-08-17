import React from "react";
const FormattedDate = ({ isoDateString }) => {
  const date = new Date(isoDateString);

  const options = {
    year: 'numeric',
    month: 'long',
    day: 'numeric',
    hour: '2-digit',
    minute: '2-digit',
    second: '2-digit',
    timeZoneName: 'short',
  };

  const formatter = new Intl.DateTimeFormat(navigator.language, options);

  const formattedDate = formatter.format(date);

  return <span>{formattedDate}</span>;
};

export default FormattedDate;