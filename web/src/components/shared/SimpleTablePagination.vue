<script lang="ts" setup>
import { computed } from 'vue';
import { useI18n } from 'vue-i18n';

const props = withDefaults(
    defineProps<{
        totalRecords: number;
        page: number;
        size: number;
        itemsPerPageOptions?: number[];
        paginationTotalVisible?: number;
    }>(),
    {
        itemsPerPageOptions: () => [10, 20, 25, 50],
        paginationTotalVisible: 7
    }
);

const emit = defineEmits<{
    'update:page': [value: number];
    'update:size': [value: number];
}>();

const { t } = useI18n();

const pages = computed(() => Math.max(1, Math.ceil(Number(props.totalRecords) / props.size)));

function onSize(v: unknown) {
    emit('update:size', Number(v));
}

function onPage(v: unknown) {
    emit('update:page', Number(v));
}
</script>

<template>
    <v-row v-if="totalRecords > 0" align="center" class="my-4 px-md-15" justify="center">
        <v-col cols="12" lg="2" md="2">
            <v-select
                :model-value="size"
                :items="itemsPerPageOptions"
                :label="t('ITEMS_PER_PAGE')"
                :hint="`${t('TOTAL_RECORD')}: ${totalRecords}`"
                class="mt-md-6"
                color="primary"
                density="compact"
                persistent-hint
                variant="outlined"
                @update:model-value="onSize"
            />
        </v-col>

        <v-col cols="12" lg="6" md="6">
            <v-pagination
                v-if="pages > 1"
                :model-value="page"
                :length="pages"
                :total-visible="paginationTotalVisible"
                class="my-5"
                rounded="circle"
                @update:model-value="onPage"
            />
            <div v-else class="d-flex justify-center align-center my-5">
                <span class="text-caption text-medium-emphasis px-2">
                    {{ t('TOTAL_RECORD') }}: {{ totalRecords }}
                </span>
            </div>
        </v-col>

        <!-- Spacer column matches shared Pagination.vue layout (sort column) for alignment -->
        <v-col class="d-none d-md-flex" cols="12" lg="2" md="2" aria-hidden="true" />
    </v-row>
</template>
