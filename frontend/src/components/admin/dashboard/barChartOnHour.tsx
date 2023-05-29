import React, { useEffect, useState } from 'react';
import {
  Chart as ChartJS,
  CategoryScale,
  LinearScale,
  BarElement,
  Title,
  Tooltip,
  Legend,
  ChartEvent,
  ActiveElement,
} from 'chart.js';
import { Bar } from 'react-chartjs-2';
import { ChartOnHourData } from '.';

ChartJS.register(
  CategoryScale,
  LinearScale,
  BarElement,
  Title,
  Tooltip,
  Legend
);

type Props = {
  title: string;
  hour: number;
  chartData: ChartOnHourData[];
};

const BarChartOnHour: React.FC<Props> = (props: Props) => {
  const [amountOfJobs, setAmountOfJobs] = useState<Array<number>>([]);
  const [labels, setLabels] = useState<Array<string>>([]);

  const addZeroWhenLowwerThanTen = (timeInt: number): string => {
    if (timeInt < 10) {
      return `0${timeInt}`
    }
    return `${timeInt}`
  }
  const createLabel = () => {
    const temp :Array<string> = []
    props.chartData.forEach((value) => {
      temp.push(`${addZeroWhenLowwerThanTen(props.hour)}:${addZeroWhenLowwerThanTen(value.minute)}`)
    })
    setLabels([...temp])
  }
  const options = {
    responsive: true,
    plugins: {
      legend: {
        position: 'top' as const,
      },
      title: {
        display: true,
        text: `${props.title}h`,
      },
      tooltip: {
        callbacks: {
          label: (context: any) => {
            const labels: Array<string> = [];
            const barData = props.chartData[context.dataIndex];
            labels.push('amount of cronjob run: ' + barData.amount_of_jobs);
            barData.cronjobs.forEach((value) => {
              labels.push(`cronjob: ${value.name}, start at: ${value.start_at}, end at: ${value.end_at}`);
            });
            return labels;
          },
        },
      },
    },
  };

  const getAmountOfJobs = () => {
    const array: Array<number> = [];
    props.chartData.forEach((value) => {
      array.push(value.amount_of_jobs);
    });
    setAmountOfJobs(array);
  };
  useEffect(() => {
    getAmountOfJobs();
    createLabel()
  }, [props.chartData]);

  const data = {
    labels,
    datasets: [
      {
        label: 'Jobs run',
        data: amountOfJobs,
        backgroundColor: 'rgba(53, 162, 235, 0.5)',
      },
    ],
  };

  return (
    <Bar
      options={options}
      data={data}
    />
  );
};

export default BarChartOnHour;
