    // Function to make API call and create table
    async function fetchDataAndCreateTable() {
        try {
          const response = await fetch('http://localhost:8080/api');
          const data = await response.json();
          createTable(data);
        } catch (error) {
          console.error('Error fetching data:', error);
        }
      }
  
      // Function to create table from API response
      function createTable(response) {
        const tableContainer = document.getElementById('app');
        tableContainer.innerHTML = ''; // Clear previous content
  
        const table = document.createElement('table');
        table.className = 'min-w-full divide-y-2 divide-gray-200 bg-white text-sm';
  
        // Create table header
        const headerRow = document.createElement('thead');
        headerRow.className = 'ltr:text-left rtl:text-right';
        const headerRowContent = `
          <tr>
            <th class="whitespace-nowrap px-4 py-2 font-medium text-gray-900 text-left">ID</th>
            <th class="whitespace-nowrap px-4 py-2 font-medium text-gray-900 text-left">Service</th>
            <th class="whitespace-nowrap px-4 py-2 font-medium text-gray-900 text-left">Status</th>
            <th class="whitespace-nowrap px-4 py-2 font-medium text-gray-900 text-left">Up time</th>
          </tr>
        `;
        headerRow.innerHTML = headerRowContent;
        table.appendChild(headerRow);
  
        // Create table body
        const tbody = document.createElement('tbody');
        tbody.className = 'divide-y divide-gray-200';
        response.forEach(item => {
          const row = document.createElement('tr');
  
          // Service
          const idCell = document.createElement('td');
          idCell.className = 'whitespace-nowrap px-4 py-2 font-medium text-gray-900';
          idCell.textContent = item.id;
          row.appendChild(idCell);

          // Service
          const serviceCell = document.createElement('td');
          serviceCell.className = 'whitespace-nowrap px-4 py-2 font-medium text-gray-900';
          serviceCell.textContent = item.service_name;
          row.appendChild(serviceCell);
  
          // Status
          const statusCell = document.createElement('td');
          statusCell.className = 'whitespace-nowrap px-4 py-2 text-gray-700';
  
          if (item.status === 'Up') {
            const statusDiv = document.createElement('div');
            statusDiv.className = 'inline-flex gap-2 rounded bg-green-100 p-1 text-black-600';
            const statusSpan = document.createElement('span');
            statusSpan.className = 'text-xs font-medium';
            statusSpan.textContent = 'Up';
            statusDiv.appendChild(statusSpan);
            statusCell.appendChild(statusDiv);
          } else {
            const statusDiv = document.createElement('div');
            statusDiv.className = 'inline-flex gap-2 rounded bg-red-100 p-1 text-black-600';
            const statusSpan = document.createElement('span');
            statusSpan.className = 'text-xs font-medium';
            statusSpan.textContent = 'Down';
            statusDiv.appendChild(statusSpan);
            statusCell.appendChild(statusDiv);
          }
  
          row.appendChild(statusCell);
  
          // Up time
          const upTimeCell = document.createElement('td');
          upTimeCell.className = 'whitespace-nowrap px-4 py-2 text-gray-700';
          upTimeCell.textContent = item.up_time;
          row.appendChild(upTimeCell);
  
          tbody.appendChild(row);
        });
  
        table.appendChild(tbody);
        tableContainer.appendChild(table);
      }
  
      // Initial call to fetch data and create the table
      fetchDataAndCreateTable();
  
      // Set interval to refresh data every 2 minutes (120,000 milliseconds)
      setInterval(() => {
        fetchDataAndCreateTable();
      }, 120000);